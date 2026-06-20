package controller
import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"em_test/src/subscription/dto"
	"em_test/src/utils"
	"strings"
)

func (c *SubController) UpdateSubFull(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr) 
	if err != nil {
		
		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка получения айди"}`, http.StatusBadRequest)
	}
	var req dto.SubRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {

		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка запроса"}`, http.StatusBadRequest)
		return
	}
	err = c.service.UpdateSubFull(id, req)

	if err != nil {
		fmt.Println("Ошибка обновления данных", err)

		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка обноваления данных"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println("Подписка успешно обновлена")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Подписка успешно обновлена."})

}

func (c *SubController) UpdateSubPartial(w http.ResponseWriter, r * http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {

		utils.LogError(err)
		http.Error(w, `{"error", "Неверный формат айди"}`, http.StatusBadRequest)
		return
	}
	var req dto.SubUpdateRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {

		utils.LogError(err)
		http.Error(w, `{"error": "Неверный JSON"}`, http.StatusBadRequest)
  	return
	}

	if ( req.EndDate == nil && req.PriceMonth == nil && req.ServiceName == nil &&
			 req.StartDate == nil && req.UserUUID == nil) {
				http.Error(w, `{"error": "Все поля пустые"}`, http.StatusBadRequest)
  			return
			}
	err = c.service.UpdateSubPartial(id, req)
	if err != nil {
	if strings.Contains(err.Error(), "не найдена") {
            http.Error(w, `{"error": "Подписка не найдена"}`, http.StatusNotFound)
            return
        }

				utils.LogError(err)
        http.Error(w, `{"error": "Ошибка обновления"}`, http.StatusInternalServerError)
        return
			}

	w.Header().Set("Content-Type", "appliaction/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
  "status":  "success", "message": "Подписка обновлена",
	})

}
