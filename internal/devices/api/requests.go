package devicesapi

type CreateDeviceRequest struct {
	UID    string `json:"uid"`
	Serial string `json:"serial"`
}
