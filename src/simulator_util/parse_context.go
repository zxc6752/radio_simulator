package simulator_util

import (
	"fmt"
	"os/exec"
	"radio_simulator/lib/ngap/ngapConvert"
	"radio_simulator/lib/openapi/models"
	"radio_simulator/lib/path_util"
	"radio_simulator/src/factory"
	"radio_simulator/src/logger"
	"radio_simulator/src/simulator_context"
	"radio_simulator/src/simulator_nas/nas_security"
	"radio_simulator/src/ue_factory"
	"strconv"
)

var self *simulator_context.Simulator = simulator_context.Simulator_Self()

func ParseRanContext() {
	config := factory.SimConfig
	self.DefaultRanSctpUri = config.RanInfo[0].RanSctpUri
	for _, ranInfo := range config.RanInfo {
		plmnId := ngapConvert.PlmnIdToNgap(ranInfo.GnbId.PlmnId)
		ran := self.AddRanContext(ranInfo.AmfUri, ranInfo.RanSctpUri, ranInfo.RanName, ranInfo.RanGtpUri, plmnId, ranInfo.GnbId.Value, ranInfo.GnbId.BitLength)
		for _, upfUri := range ranInfo.UpfUriList {
			ran.UpfInfoList[upfUri.IP] = &simulator_context.UpfInfo{
				Addr: upfUri,
			}
		}
		for _, supportItem := range ranInfo.SupportTAList {
			plmnList := []simulator_context.PlmnSupportItem{}
			for _, item := range supportItem.Plmnlist {
				plmnItem := simulator_context.PlmnSupportItem{}
				plmnItem.PlmnId = ngapConvert.PlmnIdToNgap(item.PlmnId)
				for _, snssai := range item.SNssaiList {
					sNssaiNgap := ngapConvert.SNssaiToNgap(snssai)
					plmnItem.SNssaiList = append(plmnItem.SNssaiList, sNssaiNgap)
				}
				plmnList = append(plmnList, plmnItem)
			}
			tac := TACConfigToHexString(supportItem.Tac)
			ran.DefaultTAC = tac
			ran.SupportTAList[tac] = plmnList
		}
	}
}
func ParseTunDev() {
	info := factory.SimConfig.TunnelInfo
	self.Gtp5gTunnelExec = "./" + path_util.Gofree5gcPath(info.Gtp5gPath+"/gtp5g-tunnel")
	self.TunDev = info.TunDev
	// Add default far=1, action=2(allow)
	_, err := exec.Command(self.Gtp5gTunnelExec, "add", "far", self.TunDev, "1", "--action", "2").Output()
	// iface, err := net.InterfaceByName(config.TunDev)
	if err != nil {
		logger.UtilLog.Errorf("Init default forward FAR(1) failed[%s]", err.Error())
		// logger.UtilLog.Errorf("Get Interface by Name[%s] Error[%s]", config.TunDev, err.Error())
		return
	}
	// addrs, err := iface.Addrs()
	// if err != nil {
	// 	logger.UtilLog.Errorf("Get Interface[%s] Addr Error[%s]", config.TunDev, err.Error())
	// 	return
	// }
	// uri := strings.Split(addrs[0].String(), "/")
	// var ipAddr [4]byte
	// copy(ipAddr[:], net.ParseIP(uri[0]).To4())
	// self.TunSockAddr = &syscall.SockaddrInet4{
	// 	Port: 0,
	// 	Addr: ipAddr,
	// }
	// self.TunFd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	// if err != nil {
	// 	logger.UtilLog.Error(err.Error())
	// 	return
	// }
}

func ParseUeData(configDirPath string, fileList []string) {
	for _, ueInfoFile := range fileList {
		fileName := configDirPath + ueInfoFile
		ue := ue_factory.InitUeContextFactory(fileName)
		ue.IntAlg = nas_security.AlogMaps[ue.IntegrityAlgOrig]
		ue.EncAlg = nas_security.AlogMaps[ue.CipheringAlgOrig]

		self.UeContextPool[ue.Supi] = ue
	}
}
func InitUeToDB() {
	for supi, ue := range self.UeContextPool {
		amDate := models.AccessAndMobilitySubscriptionData{
			Gpsis: ue.Gpsis,
			Nssai: &ue.Nssai,
		}
		amPolicy := models.AmPolicyData{
			SubscCats: ue.SubscCats,
		}
		auths := ue.AuthData
		authsSubs := models.AuthenticationSubscription{
			AuthenticationMethod:          models.AuthMethod(auths.AuthMethod),
			AuthenticationManagementField: auths.AMF,
			PermanentKey: &models.PermanentKey{
				PermanentKeyValue: auths.K,
			},
			SequenceNumber: auths.SQN,
		}
		if auths.Opc != "" {
			authsSubs.Opc = &models.Opc{
				OpcValue: auths.Opc,
			}
		} else if auths.Op != "" {
			authsSubs.Milenage = &models.Milenage{
				Op: &models.Op{
					OpValue: auths.Op,
				},
			}
		} else {
			logger.UtilLog.Errorf("Ue[%s] need Op or OpCode", ue.Supi)
		}
		InsertAuthSubscriptionToMongoDB(supi, authsSubs)
		InsertAccessAndMobilitySubscriptionDataToMongoDB(supi, amDate, ue.ServingPlmnId)
		InsertSmfSelectionSubscriptionDataToMongoDB(supi, ue.SmfSelData, ue.ServingPlmnId)
		InsertAmPolicyDataToMongoDB(supi, amPolicy)
	}
}

func ClearDB() {
	for supi, ue := range self.UeContextPool {
		logger.UtilLog.Infof("Del UE[%s] Info in DB", supi)
		DelAccessAndMobilitySubscriptionDataFromMongoDB(supi, ue.ServingPlmnId)
		DelAmPolicyDataFromMongoDB(supi)
		DelAuthSubscriptionToMongoDB(supi)
		DelSmfSelectionSubscriptionDataFromMongoDB(supi, ue.ServingPlmnId)
	}
}

func TACConfigToHexString(intString string) (hexString string) {
	tmp, _ := strconv.ParseUint(intString, 10, 32)
	hexString = fmt.Sprintf("%06x", tmp)
	return
}
