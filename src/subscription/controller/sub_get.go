package controller

import (
	"em_test/src/subscription/dto"
	"encoding/json"
	"fmt"
	"em_test/src/utils"
	"net/http"
	"strconv"
)

// GetSubId возвращает подписку по ID
// @Summary      Получить подписку по ID
// @Description  Возвращает данные подписки по указанному идентификатору
// @Tags         Subscriptions
// @Produce      json
// @Param        id path int true "ID подписки"
// @Success      200 {object} dto.SubRequest "Данные подписки"
// @Failure      400 "Неверный формат ID"
// @Failure      404 "Подписка не найдена"
// @Router       /get-sub/{id} [get]
func (c *SubController) GetSubId(w http.ResponseWriter, r *http.Request) {
	SubIDStr := r.PathValue("id")
	intSubID, err := strconv.Atoi(SubIDStr)

	if err != nil {
		fmt.Println("Ошибка получения айди", err)
		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка получения айди"}`, http.StatusInternalServerError)
		return
	}
	
	sub, err := c.service.GetSubId(intSubID)

	if err != nil {
		fmt.Println("Подписка не найдена", err)
		http.Error(w, `{"error": "Подписка не найдена"}`, http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(sub)


}

// SelectSubAll возвращает все подписки
// @Summary      Получить все подписки
// @Description  Возвращает список всех подписок
// @Tags         Subscriptions
// @Produce      json
// SelectSubAll
// @Success      200 {array} dto.SubRequest "Список подписок"
// @Failure      500 "Ошибка получения данных"
// @Router       /get-sub [get]
func (c *SubController) SelectSubAll(w http.ResponseWriter, r *http.Request) {
	subs, err := c.service.SelectSubAll()
	if err != nil {
		fmt.Println("Ошибка сервера", err)
		utils.LogError(err)
		http.Error(w, `"error": "Ошибка получения подписки"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subs)
}

// SelectGetSumSub возвращает сумму подписок по фильтрам
// @Summary      Получить сумму подписок
// @Description  Рассчитывает общую стоимость подписок по заданным фильтрам
// @Tags         Subscriptions
// @Produce      json
// @Param        user_id query string false "UUID пользователя"
// @Param        service_name query string false "Название сервиса"
// @Param        start_date query string true "Начальная дата в формате MM-YYYY" example(05-2026)
// @Param        end_date query string true "Конечная дата в формате MM-YYYY" example(08-2026)
// @Success      200 "Сумма подписок"
// @Failure      400 "Отсутствуют обязательные параметры"
// @Failure      500 "Ошибка получения суммы"
// @Router       /get-sub-sum [get]
func (c *SubController) SelectGetSumSub(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	serviceName := r.URL.Query().Get("service_name")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
 if startDate == "" || endDate == "" {
        http.Error(w, `{"error": "Параметры start_month и end_month обязательны"}`, http.StatusBadRequest)
        return
	}	
  filter := dto.TotalSumRequest{
        UserID:      userId,
        ServiceName: serviceName,
        DateStart:  startDate,
        DateEnd:    endDate,
    }

	cost, err := c.service.TotalGetSum(filter)

	if err != nil {
		fmt.Println("Ошибка получения суммы подписок", err)
		
		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка получения суммы"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"total_sum": cost,
		"currency": "RUB",
		"period": startDate + " - " + endDate,

	})
	
}
