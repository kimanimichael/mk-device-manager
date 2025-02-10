package devices

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kimanimichael/mk-device-manager/internal/adapters/database/sqlc/gensql"
	"time"
)

type DeviceRepositorySQL struct {
	DB *sqlcdatabase.Queries
}

var _ DeviceRepository = (*DeviceRepositorySQL)(nil)

func NewDeviceRepositorySQL(db *sqlcdatabase.Queries) *DeviceRepositorySQL {
	return &DeviceRepositorySQL{
		DB: db,
	}
}

func (r *DeviceRepositorySQL) CreateDevice(ctx context.Context, UID string, serial string) (*Device, error) {
	device, err := r.DB.CreateDevice(ctx, sqlcdatabase.CreateDeviceParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Uid:       UID,
		Serial:    serial,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %v", err)
	}
	return &Device{
		ID:        device.ID.String(),
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.CreatedAt,
		UID:       device.Uid,
		Serial:    device.Serial,
	}, nil
}

func (r *DeviceRepositorySQL) GetDeviceByID(ctx context.Context, ID string) (*Device, error) {
	deviceID, err := uuid.Parse(ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse device UUID: %v", err)
	}
	device, err := r.DB.GetDeviceByID(ctx, deviceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get device by UUID: %v", err)
	}
	return &Device{
		ID:        device.ID.String(),
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.CreatedAt,
		UID:       device.Uid,
		Serial:    device.Serial,
	}, nil
}

func (r *DeviceRepositorySQL) GetDeviceByUID(ctx context.Context, UID string) (*Device, error) {
	device, err := r.DB.GetDeviceByUID(ctx, UID)
	if err != nil {
		return nil, fmt.Errorf("failed to get device by UUID: %v", err)
	}
	return &Device{
		ID:        device.ID.String(),
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.CreatedAt,
		UID:       device.Uid,
		Serial:    device.Serial,
	}, nil
}

func (r *DeviceRepositorySQL) GetDeviceBySerial(ctx context.Context, Serial string) (*Device, error) {
	device, err := r.DB.GetDeviceBySerial(ctx, Serial)
	if err != nil {
		return nil, fmt.Errorf("failed to get device by Serial: %v", err)
	}
	return &Device{
		ID:        device.ID.String(),
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.CreatedAt,
		UID:       device.Uid,
		Serial:    device.Serial,
	}, nil
}

func (r *DeviceRepositorySQL) GetDevices(ctx context.Context) ([]Device, error) {
	devices, err := r.DB.GetDevices(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get device by Uerial: %v", err)
	}
	var devicesToReturn []Device
	for _, device := range devices {
		devicesToReturn = append(devicesToReturn, Device{
			ID:        device.ID.String(),
			CreatedAt: device.CreatedAt,
			UpdatedAt: device.CreatedAt,
			UID:       device.Uid,
			Serial:    device.Serial,
		})
	}
	return devicesToReturn, nil
}

func (r *DeviceRepositorySQL) GetPagedDevices(ctx context.Context, offset, limit uint32) (*DevicePage, error) {
	devices, err := r.DB.GetPagedDevices(ctx, sqlcdatabase.GetPagedDevicesParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get paged devices: %v", err)
	}
	var devicesToReturn []Device
	for _, device := range devices {
		devicesToReturn = append(devicesToReturn, Device{
			ID:        device.ID.String(),
			CreatedAt: device.CreatedAt,
			UpdatedAt: device.CreatedAt,
			UID:       device.Uid,
			Serial:    device.Serial,
		})
	}
	totalDevices, err := r.DB.GetDevicesCount(ctx)
	devicePage := DevicePage{
		page: page{
			Offset: offset,
			Total:  uint32(totalDevices),
		},
		Devices: devicesToReturn,
	}
	return &devicePage, nil
}

func (r *DeviceRepositorySQL) DeleteDevice(ctx context.Context, ID string) error {
	deviceID, err := uuid.Parse(ID)
	if err != nil {
		return fmt.Errorf("failed to parse device UUID: %v", err)
	}

	_, err = r.DB.GetDeviceByID(ctx, deviceID)
	if err != nil {
		return fmt.Errorf("error getting device from ID: %v", err)
	}

	err = r.DB.DeleteDevice(ctx, deviceID)
	if err != nil {
		return fmt.Errorf("failed to delete device: %v", err)
	}
	return nil
}
