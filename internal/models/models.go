package models

import "time"

// AnalyticsEvent represents analytics event types
type AnalyticsEvent string

const (
	UE_MOBILITY         AnalyticsEvent = "UE_MOBILITY"
	UE_COMM             AnalyticsEvent = "UE_COMM"
	ABNORMAL_BEHAVIOR   AnalyticsEvent = "ABNORMAL_BEHAVIOR"
	CONGESTION          AnalyticsEvent = "CONGESTION"
	NETWORK_PERFORMANCE AnalyticsEvent = "NETWORK_PERFORMANCE"
	QOS_SUSTAINABILITY  AnalyticsEvent = "QOS_SUSTAINABILITY"
)

// AnalyticsExposureSubsc represents a subscription to analytics exposure
type AnalyticsExposureSubsc struct {
	AnalyticsEventSubsc []AnalyticsEventSubsc `json:"analyEventsSubs"`
	AnalyticsRepInfo    *ReportingInfo        `json:"analyRepInfo,omitempty"`
	NotifURI            string                `json:"notifUri"`
	NotifID             string                `json:"notifId"`
	EventNotifications  []AnalyticsEventNotif `json:"eventNotifis,omitempty"`
	SuppFeat            *string               `json:"suppFeat,omitempty"`
}

// AnalyticsEventSubsc represents a subscription to a specific analytics event
type AnalyticsEventSubsc struct {
	AnalyticsEvent AnalyticsEvent             `json:"analyEvent"`
	EventFilter    *AnalyticsEventFilterSubsc `json:"analyEventFilter,omitempty"`
	TargetUE       *TargetUeID                `json:"tgtUe,omitempty"`
}

// AnalyticsEventFilterSubsc represents filters for analytics event subscription
type AnalyticsEventFilterSubsc struct {
	NetworkPerfReqs []string `json:"nwPerfReqs,omitempty"`
	LocationArea    *string  `json:"locArea,omitempty"`
	ApplicationIDs  []string `json:"appIds,omitempty"`
	SNSSAI          *string  `json:"snssai,omitempty"`
}

// TargetUeID represents target UE identification
type TargetUeID struct {
	AnyUEIndication bool    `json:"anyUeInd,omitempty"`
	GPSI            *string `json:"gpsi,omitempty"`
	ExternalGroupID *string `json:"exterGroupId,omitempty"`
}

// AnalyticsEventNotif represents a notification for an analytics event
type AnalyticsEventNotif struct {
	AnalyticsEvent      AnalyticsEvent    `json:"analyEvent"`
	Timestamp           time.Time         `json:"timeStamp"`
	UEMobilityInfo      []UEMobilityInfo  `json:"ueMobilityInfos,omitempty"`
	UECommunicationInfo []string          `json:"ueCommInfos,omitempty"`
	AbnormalInfo        []string          `json:"abnormalInfos,omitempty"`
	CongestionInfo      []string          `json:"congestInfos,omitempty"`
	NetworkPerfInfo     []NetworkPerfInfo `json:"nwPerfInfos,omitempty"`
	QoSSustainInfo      []string          `json:"qosSustainInfos,omitempty"`
}

// AnalyticsEventNotification is the complete notification structure
type AnalyticsEventNotification struct {
	NotifID               string                `json:"notifId"`
	AnalyticsEventNotifis []AnalyticsEventNotif `json:"analyEventNotifs"`
}

// UEMobilityInfo represents UE mobility exposure information
type UEMobilityInfo struct {
	Timestamp        time.Time `json:"ts,omitempty"`
	Duration         uint32    `json:"duration"`
	DurationVariance float32   `json:"durationVariance,omitempty"`
	LocationInfo     []string  `json:"locInfo"`
	Ratio            float32   `json:"ratio,omitempty"`
	Confidence       uint32    `json:"confidence,omitempty"`
}

// NetworkPerfInfo represents network performance information
type NetworkPerfInfo struct {
	LocationArea    string  `json:"locArea"`
	NetworkPerfType string  `json:"nwPerfType"`
	RelativeRatio   float32 `json:"relativeRatio,omitempty"`
	AbsoluteNum     uint32  `json:"absoluteNum,omitempty"`
	Confidence      uint32  `json:"confidence,omitempty"`
}

// ReportingInfo represents reporting information
type ReportingInfo struct {
	MaxReportNbr uint32 `json:"maxRepNbr,omitempty"`
	Interval     uint32 `json:"interval,omitempty"`
}

// AnalyticsRequest represents a request for analytics data
type AnalyticsRequest struct {
	AnalyticsEvent AnalyticsEvent        `json:"analyEvent"`
	EventFilter    *AnalyticsEventFilter `json:"analyEventFilter,omitempty"`
	ReportingInfo  *ReportingInfo        `json:"analyRep,omitempty"`
	TargetUE       *TargetUeID           `json:"tgtUe,omitempty"`
	SuppFeat       string                `json:"suppFeat"`
}

// AnalyticsEventFilter represents filters for analytics data fetch requests
type AnalyticsEventFilter struct {
	LocationArea     *string  `json:"locArea,omitempty"`
	DNN              *string  `json:"dnn,omitempty"`
	NetworkPerfTypes []string `json:"nwPerfTypes,omitempty"`
	ApplicationIDs   []string `json:"appIds,omitempty"`
	ExceptionIDs     []string `json:"excepIds,omitempty"`
	SNSSAI           *string  `json:"snssai,omitempty"`
	QoSRequirement   *string  `json:"qosReq,omitempty"`
}

// AnalyticsData represents the response containing analytics data
type AnalyticsData struct {
	UEMobilityInfos      []UEMobilityInfo  `json:"ueMobilityInfos,omitempty"`
	UECommunicationInfos []string          `json:"ueCommInfos,omitempty"`
	NetworkPerfInfos     []NetworkPerfInfo `json:"nwPerfInfos,omitempty"`
	AbnormalInfos        []string          `json:"abnormalInfos,omitempty"`
	CongestionInfos      []string          `json:"congestInfos,omitempty"`
	QoSSustainInfos      []string          `json:"qosSustainInfos,omitempty"`
	SuppFeat             *string           `json:"suppFeat,omitempty"`
}
