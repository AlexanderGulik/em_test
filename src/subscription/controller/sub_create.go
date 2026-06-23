package controller
import (
	"net/http"
	"em_test/src/subscription/dto"
	"em_test/src/utils"
	"encoding/json"
	"fmt"
	"time"
)

// CreateSub создает новую подписку
// @Summary      Создать подписку
// @Description  Создает новую подписку с указанными параметрами
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Param        request body dto.SubRequest true "Данные для создания подписки"
// @Success      201 "Подписка успешно создана"
// @Failure      400 "Ошибка валидации данных"
// @Failure      500 "Внутренняя ошибка сервера"
// @Router       /created-sub [post]
func (c *SubController) CreateSub(w http.ResponseWriter, r *http.Request){

	var req dto.SubRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка декодирования запроса"}`, http.StatusBadRequest)
		return
	}
	if !isValidDateRange(req.DateStart, req.DateEnd){
		http.Error(w, `{"error": "Ошибка формата даты"}`, http.StatusBadRequest)
		return

	}
	
	 id, err := c.service.CreateSub(req);
	if err != nil {
		fmt.Printf("Ошибка создания: %v", err)
		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка сохранения"}`, http.StatusInternalServerError)
		return
	}
  response := dto.CreateSubscriptionResponse{
        Status:  "success",
        Message: "Подписка успешно добавлена.",
        ID:      id,
    }
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Успешно добавлена новая подписка")
	json.NewEncoder(w).Encode(response)

}

func isValidDateRange(dateStart, dateEnd string) bool {
    start, err := time.Parse("01-2006", dateStart)
    if err != nil {
        return false
    }

    end, err := time.Parse("01-2006", dateEnd)
    if err != nil {
        return false
    }

    return start.Before(end)
}





