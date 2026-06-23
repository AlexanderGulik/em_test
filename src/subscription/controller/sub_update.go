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

// UpdateSubFull полностью обновляет подписку
// @Summary      Полное обновление подписки
// @Description  Заменяет все поля подписки по указанному ID
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Param        id path int true "ID подписки"
// @Param        request body dto.SubRequest true "Новые данные подписки"
// @Success      200 "Подписка успешно обновлена"
// @Failure      400 "Ошибка валидации данных"
// @Failure      404 "Подписка не найдена"
// @Failure      500 "Ошибка обновления"
// @Router       /update-sub/{id} [put]
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
		if strings.Contains(err.Error(),"не найдена") {
			fmt.Println("Подписка не найдена", id)
			http.Error(w,`{"error": "Подписка не найдена"}`, http.StatusNotFound)
			return
		}
		fmt.Println("Ошибка обновления данных", err)

		utils.LogError(err)
		http.Error(w, `{"error": "Ошибка обноваления данных"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println("Подписка успешно обновлена")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Подписка успешно обновлена."})

}

// UpdateSubPartial частично обновляет подписку
// @Summary      Частичное обновление подписки
// @Description  Обновляет только указанные поля подписки
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Param        id path int true "ID подписки"
// @Param        request body dto.SubUpdateRequest true "Поля для обновления"
// @Success      200 "Подписка обновлена"
// @Failure      400 "Ошибка валидации или все поля пустые"
// @Failure      404 "Подписка не найдена"
// @Failure      500 "Ошибка обновления"
// @Router       /update-sub/{id} [patch]
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
  "status":  "success", "message": "Подписка обновлена",
	})

}
