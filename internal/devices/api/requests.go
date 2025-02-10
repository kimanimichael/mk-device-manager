package devicesapi

type CreateDeviceRequest struct {
	UID    string `json:"uid"`
	Serial string `json:"serial"`
}

type GetDeviceByIDRequest struct {
	ID string `json:"id"`
}

type GetDeviceByUIDRequest struct {
	UID string `json:"device_uid"`
}

type GetDeviceBySerialRequest struct {
	Serial string `json:"device_serial"`
}

type GetPagedDevicesRequest struct {
	Offset uint32 `json:"offset"`
	Limit  uint32 `json:"limit"`
}
