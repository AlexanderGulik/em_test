package controller
import (
	"net/http"
	"em_test/src/subscription/dto"
	"em_test/src/utils"
	"encoding/json"
	"fmt"
	"time"
)

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
	
	 err := c.service.CreateSub(req);
	if err != nil {
		fmt.Printf("Ошибка создания: %v", err)
		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка сохранения"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Println("Успешно добавлена новая подписка")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Подписка успешно добавлена."})

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





