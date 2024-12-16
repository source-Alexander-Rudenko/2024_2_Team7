package events

import (
	"context"
	"net/http"
	"strconv"
	"time"

	pbEvent "kudago/internal/event/api"
	httpErrors "kudago/internal/gateway/errors"
	"kudago/internal/gateway/utils"
	pbNtf "kudago/internal/notification/api"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	grpcCodes "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

// UpdateEvent обновляет данные существующего события.
// @Summary Обновление события
// @Description Обновляет данные существующего события. Необходимо передать JSON-объект с данными события и идентификатором события в URL.
// @Tags events
// @Accept  json
// @Produce  json
// @Param id path int true "Идентификатор события"
// @Param json body NewEventRequest true "Данные для обновления события"
// @Param image formData file false "Изображение события"
// @Success 200 {object} NewEventResponse "Успешное обновление события"
// @Failure 400 {object} httpErrors.HttpError "Неверные данные"
// @Failure 401 {object} httpErrors.HttpError "Неавторизован"
// @Failure 403 {object} httpErrors.HttpError "Доступ запрещен"
// @Failure 404 {object} httpErrors.HttpError "Событие не найдено"
// @Failure 500 {object} httpErrors.HttpError "Внутренняя ошибка сервера"
// @Router /events/{id} [put]
func (h EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	session, ok := utils.GetSessionFromContext(r.Context())
	if !ok {
		utils.WriteResponse(w, http.StatusForbidden, httpErrors.ErrUnauthorized)
		return
	}

	req, media, reqErr := parseEventData(r)
	if reqErr != nil {
		utils.WriteResponse(w, http.StatusBadRequest, reqErr)
		return
	}

	_, err := govalidator.ValidateStruct(&req)
	if err != nil {
		utils.ProcessValidationErrors(w, err)
		return
	}

	reqErr = checkNewEventRequest(req)
	if reqErr != nil {
		utils.WriteResponse(w, http.StatusBadRequest, reqErr)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, httpErrors.ErrInvalidID)
		return
	}

	url, err := h.uploadImage(r.Context(), media, w)
	if err != nil {
		return
	}

	event := toPBEvent(req, session.UserID)
	event.ID = int32(id)
	event.Image = url

	event, err = h.EventService.UpdateEvent(r.Context(), event)
	if err != nil {
		h.deleteImage(r.Context(), url)
		st, ok := grpcStatus.FromError(err)
		if ok {
			switch st.Code() {
			case grpcCodes.NotFound:
				utils.WriteResponse(w, http.StatusConflict, httpErrors.ErrEventNotFound)
				return
			case grpcCodes.PermissionDenied:
				utils.WriteResponse(w, http.StatusForbidden, httpErrors.ErrAccessDenied)
				return
			case grpcCodes.Internal:
				utils.WriteResponse(w, http.StatusInternalServerError, httpErrors.ErrInternal)
				return
			}
		}

		h.logger.Error(r.Context(), "update event", err)
		utils.WriteResponse(w, http.StatusInternalServerError, httpErrors.ErrInternal)
		return
	}

	err = h.sendUpdateNotifications(r.Context(), int(event.ID))
	if err != nil {
		h.logger.Error(r.Context(), "send update notifications", err)
	}

	eventResp := eventToEventResponse(event)
	resp := NewEventResponse{
		Event: eventResp,
	}
	utils.WriteResponse(w, http.StatusOK, resp)
}

func (h EventHandler) sendUpdateNotifications(ctx context.Context, eventID int) error {
	idsResp, err := h.EventService.GetUserIDsByFavoriteEvent(ctx, &pbEvent.GetUserIDsByFavoriteEventRequest{ID: int32(eventID)})
	if err != nil {
		return err
	}

	req := &pbNtf.CreateNotificationsRequest{
		UserIDs: make([]int32, len(idsResp.IDs)),
		Notification: &pbNtf.Notification{
			Message:  UpdatedEventMsg,
			NotifyAt: time.Now().String(),
			EventID:  int32(eventID),
		},
	}

	for _, id := range idsResp.IDs {
		req.UserIDs = append(req.UserIDs, int32(id))
	}

	_, err = h.NotificationService.CreateNotifications(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
