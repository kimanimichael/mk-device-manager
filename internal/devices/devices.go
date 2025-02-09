package devices

import (
	"context"
	"time"
)

type Device struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	UID       string
	Serial    string
}

type page struct {
	Offset uint32
	Total  uint32
}

type DevicePage struct {
	page
	Devices []Device
}

type DeviceService interface {
	CreateDevice(ctx context.Context, UID string, serial string) (*Device, error)
	GetDeviceByID(ctx context.Context, ID string) (*Device, error)
	GetDeviceByUID(ctx context.Context, UID string) (*Device, error)
	GetDeviceBySerial(ctx context.Context, Serial string) (*Device, error)
	GetDevices(ctx context.Context) ([]Device, error)
	GetPagedDevices(ctx context.Context, offset uint32, limit uint32) (*DevicePage, error)
	DeleteDevice(ctx context.Context, ID string) error
}

type DeviceRepository interface {
	CreateDevice(ctx context.Context, UID string, serial string) (*Device, error)
	GetDeviceByID(ctx context.Context, ID string) (*Device, error)
	GetDeviceByUID(ctx context.Context, UID string) (*Device, error)
	GetDeviceBySerial(ctx context.Context, Serial string) (*Device, error)
	GetDevices(ctx context.Context) ([]Device, error)
	GetPagedDevices(ctx context.Context, offset uint32, limit uint32) (*DevicePage, error)
	DeleteDevice(ctx context.Context, ID string) error
}
