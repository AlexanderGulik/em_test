package controller

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"em_test/src/utils"
)

// DeleteSub удаляет подписку по ID
// @Summary      Удалить подписку
// @Description  Удаляет подписку по указанному ID
// @Tags         Subscriptions
// @Produce      json
// @Param        id path int true "ID подписки"
// @Success      200 "Подписка успешно удалена"
// @Failure      400 "Неверный формат ID"
// @Failure      500 "Ошибка удаления"
// @Router       /delete-sub/{id} [delete]
func (c *SubController) DeleteSub(w http.ResponseWriter, r *http.Request) {
	subIDStr := r.PathValue("id")
	intSubId, err := strconv.Atoi(subIDStr)
	if err != nil {
		fmt.Println("Ошибка получения айди", err)
		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка удаленя"}`, http.StatusInternalServerError)
		return

	}

	err = c.service.DeleteSubId(intSubId)

	if err != nil {

		utils.LogError(err)
		fmt.Println("Ошибка удаления подписки", http.StatusInternalServerError)
		return
	}
	fmt.Println("Подписка успешно удалена")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Подписка успешно удалена."})


}
