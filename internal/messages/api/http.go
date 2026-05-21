package messagesapi

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/kimanimichael/mk-device-manager/internal/devices"
	"github.com/kimanimichael/mk-device-manager/internal/messages"
	"github.com/mike-kimani/fechronizo/v2/pkg/httpresponses"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	defaultOffset uint32 = 0
	defaultLimit  uint32 = 20
)

func parsePagination(r *http.Request) (uint32, uint32, error) {
	offset := defaultOffset
	limit := defaultLimit

	if v := r.URL.Query().Get("offset"); v != "" {
		parsed, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, 0, err
		}
		offset = uint32(parsed)
	}

	if v := r.URL.Query().Get("limit"); v != "" {
		parsed, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, 0, err
		}
		limit = uint32(parsed)
	}

	return offset, limit, nil
}

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
	id := r.URL.Query().Get("id")
	if id == "" {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "id query parameter is required")
		return
	}

	ctx := r.Context()

	message, err := h.service.GetMessageByID(ctx, id)
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

	offset, limit, err := parsePagination(r)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Invalid offset or limit query parameter")
		return
	}

	ctx := r.Context()
	_, err = h.deviceService.GetDeviceByUID(ctx, uid)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	returnedMessages, err := h.service.GetMessagesByUID(ctx, uid, offset, limit)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, returnedMessages)
}

func (h *MessageHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	offset, limit, err := parsePagination(r)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusBadRequest, "Invalid offset or limit query parameter")
		return
	}

	ctx := r.Context()
	returnedMessages, err := h.service.GetAllMessages(ctx, offset, limit)
	if err != nil {
		httpresponses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpresponses.RespondWithJson(w, http.StatusOK, returnedMessages)
}
