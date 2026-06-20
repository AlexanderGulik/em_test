package service

import (
	"fmt"
	"em_test/src/subscription/dto"
	"em_test/src/config"
	"strings"
)



func (s *SubService) UpdateSubFull(id int, req dto.SubRequest) (error) {
	_, err := config.DB.Exec(`UPDATE subscriptions SET service_name = $1, price_month = $2, user_uuid = $3, start_date = $4, end_date = $5 WHERE id_sub = $6`,
			req.NameSub, req.PriceMonth, req.UserId, req.DateStart, req.DateEnd, id)
		if err != nil {
		fmt.Println("Ошибка выполнения запроса")
		return  err
	}
	return nil

}
func (s *SubService) UpdateSubPartial(id int, req dto.SubUpdateRequest) error {
	fmt.Println(req)
    updates := map[string]interface{}{}

    if req.ServiceName != nil {
        updates["service_name"] = *req.ServiceName
    }
    if req.PriceMonth != nil {
        updates["price_month"] = *req.PriceMonth
    }
    if req.UserUUID != nil {
        updates["user_uuid"] = *req.UserUUID
    }
    if req.StartDate != nil {
        updates["start_date"] = *req.StartDate
    }
    if req.EndDate != nil {
        updates["end_date"] = *req.EndDate
    }

    if len(updates) == 0 {
        return nil
    }

    updates["id_sub"] = id

   	var sets []string
    for field := range updates {
        sets = append(sets, fmt.Sprintf("%s = :%s", field, field))
    }
    query := fmt.Sprintf("UPDATE subscriptions SET %s WHERE id_sub = :id_sub", strings.Join(sets, ", "))

    updates["id_sub"] = id
    _, err := config.DB.NamedExec(query, updates)
    return err
	}
