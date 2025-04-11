package messagesapi

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/kimanimichael/mk-device-manager/internal/devices"
	"github.com/kimanimichael/mk-device-manager/internal/messages"
	"github.com/mike-kimani/fechronizo/v2/pkg/httpresponses"
	"io/ioutil"
	"net/http"
)

type MessageHandler struct {
	service       messages.MessageService
	deviceService devices.DeviceService
}

func NewMessageHandler(service messages.MessageService, deviceService devices.DeviceService) *MessageHandler {
	return &MessageHandler{
		service:       service,
		deviceService: deviceService,
	}
}

func (h *MessageHandler) RegisterRoutes(router chi.Router) {
	router.Post("/message/{uid}", h.CreateMessage)
	router.Get("/message", h.GetMessageByID)
	router.Get("/messages/{uid}", h.GetDeviceMessagesByUID)
	router.Get("/messages", h.GetAllMessages)
}

func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	if uid == "" {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Missing UID")
		return
	}

	ctx := r.Context()
	_, err := h.deviceService.GetDeviceByUID(ctx, uid)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var js json.RawMessage
	if err := json.Unmarshal(rawBody, &js); err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return
	}

	msg := &messages.Message{
		DeviceUID: uid,
		Payload:   js,
	}

	message, err := h.service.CreateMessage(ctx, msg)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, messageToMessageResponse(*message))
}

func (h *MessageHandler) GetMessageByID(w http.ResponseWriter, r *http.Request) {
	params := GetMessageByIDRequest{}

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

	message, err := h.service.GetMessageByID(ctx, params.ID)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, messageToMessageResponse(*message))
}

func (h *MessageHandler) GetDeviceMessagesByUID(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	if uid == "" {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Missing UID")
		return
	}

	params := GetMessagesByUIDRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Request body could not be decoded as JSON: %v", err))
		return
	}
	if params.Limit == 0 {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Limit field is required")
		return
	}

	ctx := r.Context()
	_, err := h.deviceService.GetDeviceByUID(ctx, uid)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	returnedMessages, err := h.service.GetMessagesByUID(ctx, uid, params.Offset, params.Limit)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, returnedMessages)
}

func (h *MessageHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	params := GetMessagesByUIDRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Request body could not be decoded as JSON: %v", err))
		return
	}
	if params.Limit == 0 {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Limit field is required")
		return
	}
	ctx := r.Context()
	returnedMessages, err := h.service.GetAllMessages(ctx, params.Offset, params.Limit)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, returnedMessages)
}
