package events

import (
	"net/http"

	httpErrors "kudago/internal/gateway/errors"
	"kudago/internal/gateway/utils"
)

// @Summary Получить все категории
// @Description Получить список всех доступных категорий событий
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Category "Список категорий"
// @Failure 500 {object} httpErrors.HttpError "Internal Server Error"
// @Router /categories [get]
func (h EventHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.EventService.GetCategories(r.Context(), nil)
	if err != nil {
		h.logger.Error(r.Context(), "get categories", err)
		utils.WriteResponse(w, http.StatusInternalServerError, httpErrors.ErrInternal)
		return
	}

	utils.WriteResponse(w, http.StatusOK, categories)
}
