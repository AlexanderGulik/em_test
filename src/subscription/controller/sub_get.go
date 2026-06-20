package controller

import (
	"em_test/src/subscription/dto"
	"encoding/json"
	"fmt"
	"em_test/src/utils"
	"net/http"
	"strconv"
)
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
