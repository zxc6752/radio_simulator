/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type NsmfEventExposure struct {
	Supi string `json:"supi,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	// Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \"FALSE\" is used, if not present.
	AnyUeInd bool `json:"anyUeInd,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	PduSeId int32 `json:"pduSeId,omitempty"`

	// Identifies an Individual SMF Notification Subscription. To enable that the value is used as part of a URI, the string shall only contain characters allowed according to the \"lower-with-hyphen\" naming convention defined in 3GPP TS 29.501 [2]. In an OpenAPI [10] schema, the format shall be designated as \"SubId\".
	SubId string `json:"subId,omitempty"`

	// Notification Correlation ID assigned by the NF service consumer.
	NotifId string `json:"notifId"`

	NotifUri string `json:"notifUri"`

	// Alternate or backup IPv4 Addess(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`

	// Alternate or backup IPv6 Addess(es) where to send Notifications.
	AltNotifIpv6Addrs []string `json:"altNotifIpv6Addrs,omitempty"`

	// Subscribed events
	EventSubs []EventSubscription `json:"eventSubs"`

	ImmeRep bool `json:"ImmeRep,omitempty"`

	NotifMethod NotificationMethod `json:"notifMethod,omitempty"`

	MaxReportNbr int32 `json:"maxReportNbr,omitempty"`

	Expiry *time.Time `json:"expiry,omitempty"`

	RepPeriod int32 `json:"repPeriod,omitempty"`

	Guami *Guami `json:"guami,omitempty"`

	// If the NF service consumer is an AMF, it should provide the name of a service produced by the AMF that makes use of notifications about subscribed events.
	ServiveName string `json:"serviveName,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
