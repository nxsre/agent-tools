package models

// MegaCliPhysicalDisk ..
type MegaCliPhysicalDisk struct {
	EncDeviceId                         int    `json:"enclosure_device_id"`
	SlotNumber                          int    `json:"slot_number"`
	Wwn                                 string `json:"wwn"`
	MedErrCount                         string `json:"media_error_count"`
	OthErrCount                         string `json:"other_error_count"`
	PredictiveFailureCount              string `json:"predictive_failure_ount"`
	LastPredictiveFailureEventSeqNumber string `json:"last_predictive_failure_event_seq_number"`
	PdType                              string `json:"pd_type"`
	RawSize                             string `json:"raw_size"`
	FirmwareState                       string `json:"firmware_state"`
	InquiryData                         string `json:"inquiry_data"`
	DeviceSpeed                         string `json:"device_speed"`
	LinkSpeed                           string `json:"link_speed"`
	MediaType                           string `json:"media_type"`
	DriveTemp                           string `json:"drive_temp"`
}
