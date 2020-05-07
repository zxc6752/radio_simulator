package simulator_message

import (
	"sync"
)

var SimChannel map[string]chan NGAPMessage
var Mtx map[string]*sync.Mutex

type NGAPMessage struct {
	Value []byte // input/request value
}

const (
	MaxChannel int = 100000
)

func init() {
	// init Pool
	SimChannel = make(map[string]chan NGAPMessage, MaxChannel)
	Mtx = make(map[string]*sync.Mutex)
}

func SendMessage(laddr string, msg NGAPMessage) {
	Mtx[laddr].Lock()
	SimChannel[laddr] <- msg
	Mtx[laddr].Unlock()
}

func AddNgapChannel(laddr string) {
	SimChannel[laddr] = make(chan NGAPMessage)
	Mtx[laddr] = &sync.Mutex{}
}
func DelNgapChannel(laddr string) {
	delete(SimChannel, laddr)
	delete(Mtx, laddr)
}
