package devicesapi

import (
	"github.com/kimanimichael/mk-device-manager/internal/devices"
	"time"
)

type DeviceResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UID       string    `json:"uid"`
	Serial    string    `json:"serial"`
}

func deviceToDeviceResponse(device devices.Device) DeviceResponse {
	return DeviceResponse{
		ID:        device.ID,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
		UID:       device.UID,
		Serial:    device.Serial,
	}
}

func devicesToDevicesResponse(devices []devices.Device) []DeviceResponse {
	var devicesResponse []DeviceResponse
	for _, device := range devices {
		devicesResponse = append(devicesResponse, deviceToDeviceResponse(device))
	}
	return devicesResponse
}
