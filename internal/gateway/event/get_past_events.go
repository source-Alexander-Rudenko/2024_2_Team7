package events

import (
	"net/http"

	httpErrors "kudago/internal/gateway/errors"
	"kudago/internal/gateway/utils"
)

// @Summary Получить все прошедшие события
// @Description Получить все прошедшие события
// @Tags events
// @Accept  json
// @Produce  json
// @Param page query int false "Номер страницы (по умолчанию 1)"
// @Param limit query int false "Количество событий на странице (по умолчанию 30)"
// @Success 200 {object} GetEventsResponse
// @Failure 500 {object} httpErrors.HttpError "Internal Server Error"
// @Router /events [get]
func (h EventHandler) GetPastEvents(w http.ResponseWriter, r *http.Request) {
	paginationParams := GetPaginationParams(r)

	events, err := h.EventService.GetPastEvents(r.Context(), paginationParams)
	if err != nil {
		h.logger.Error(r.Context(), "getPastEvents", err)
		utils.WriteResponse(w, http.StatusInternalServerError, httpErrors.ErrInternal)
		return
	}
	resp := writeEventsResponse(events.Events, int(paginationParams.Limit))

	utils.WriteResponse(w, http.StatusOK, resp)
}
