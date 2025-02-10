package devicesapi

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/kimanimichael/mk-device-manager/internal/devices"
	"github.com/mike-kimani/fechronizo/v2/pkg/httpresponses"
	"net/http"
)

type DeviceHandler struct {
	service devices.DeviceService
}

func NewDeviceHandler(service devices.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		service: service,
	}
}

func (h *DeviceHandler) RegisterRoutes(router chi.Router) {
	router.Post("/device", h.CreateDevice)
	router.Get("/device", h.GetDeviceFromID)
	router.Get("/devices", h.GetDevices)
	router.Get("/paged_devices", h.GetPagedDevices)
}

func (h *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	params := CreateDeviceRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Request body could not be decoded as JSON: %v", err))
		return
	}

	if params.UID == "" || params.Serial == "" {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Both serial and uid must be populated")
		return
	}

	if len(params.Serial) < 10 || len(params.UID) < 10 {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Each of serial and uid must be longer than 10 characters")
		return
	}

	ctx := r.Context()

	device, err := h.service.CreateDevice(ctx, params.UID, params.Serial)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusCreated, deviceToDeviceResponse(*device))
}

func (h *DeviceHandler) GetDeviceFromID(w http.ResponseWriter, r *http.Request) {
	params := GetDeviceByIDRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Request body could not be decoded as JSON: %v", err))
		return
	}

	if params.ID == "" {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "ID field is required")
		return
	}

	ctx := r.Context()

	device, err := h.service.GetDeviceByID(ctx, params.ID)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, deviceToDeviceResponse(*device))
}

func (h *DeviceHandler) GetDeviceFromUID(w http.ResponseWriter, r *http.Request) {
	params := GetDeviceByUIDRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Request body could not be decoded as JSON: %v", err))
		return
	}

	if params.UID == "" {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "UID field is required")
		return
	}

	ctx := r.Context()

	device, err := h.service.GetDeviceByUID(ctx, params.UID)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, deviceToDeviceResponse(*device))
}

func (h *DeviceHandler) GetDeviceFromSerial(w http.ResponseWriter, r *http.Request) {
	params := GetDeviceBySerialRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Request body could not be decoded as JSON: %v", err))
		return
	}

	if params.Serial == "" {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Serial field is required")
		return
	}

	ctx := r.Context()

	device, err := h.service.GetDeviceBySerial(ctx, params.Serial)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, deviceToDeviceResponse(*device))
}

func (h *DeviceHandler) GetDevices(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	devicesResponse, err := h.service.GetDevices(ctx)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpresponses.RespondWithJson(w, http.StatusOK, devicesToDevicesResponse(devicesResponse))
}

func (h *DeviceHandler) GetPagedDevices(w http.ResponseWriter, r *http.Request) {
	params := GetPagedDevicesRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("failed to decode request body"))
	}
	ctx := r.Context()
	devicesPage, err := h.service.GetPagedDevices(ctx, params.Offset, params.Limit)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	httpresponses.RespondWithJson(w, http.StatusOK, devicesPage)
}

func (h *DeviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "device_id")

	ctx := r.Context()

	err := h.service.DeleteDevice(ctx, deviceID)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	httpresponses.RespondWithJson(w, http.StatusNoContent, fmt.Sprintf("Device successfully deleted"))
}
