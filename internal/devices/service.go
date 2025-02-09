package devices

import "context"

type deviceService struct {
	repo DeviceRepository
}

func NewDeviceService(repo DeviceRepository) DeviceService {
	return &deviceService{
		repo: repo,
	}
}

func (s *deviceService) CreateDevice(ctx context.Context, UID string, serial string) (*Device, error) {
	device, err := s.repo.CreateDevice(ctx, UID, serial)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *deviceService) GetDeviceByID(ctx context.Context, ID string) (*Device, error) {
	device, err := s.repo.GetDeviceByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *deviceService) GetDeviceByUID(ctx context.Context, UID string) (*Device, error) {
	device, err := s.repo.GetDeviceByUID(ctx, UID)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *deviceService) GetDeviceBySerial(ctx context.Context, serial string) (*Device, error) {
	device, err := s.repo.GetDeviceBySerial(ctx, serial)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *deviceService) GetDevices(ctx context.Context) ([]Device, error) {
	device, err := s.repo.GetDevices(ctx)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *deviceService) GetPagedDevices(ctx context.Context, offset uint32, limit uint32) (*DevicePage, error) {
	device, err := s.repo.GetPagedDevices(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *deviceService) DeleteDevice(ctx context.Context, ID string) error {
	err := s.repo.DeleteDevice(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
