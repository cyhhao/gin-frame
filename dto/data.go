package dto

type HeaderInfo struct {
	ClientIP    string `json:"client_ip"`
	AppVersion  string `json:"app_version"`
	PhoneSystem string `json:"phone_system"`
	DeviceID    string `json:"device_id"`
}
