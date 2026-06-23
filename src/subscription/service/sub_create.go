package service
import (
	"fmt"
	"em_test/src/subscription/dto"
	"em_test/src/config"
)
func (s *SubService) CreateSub(req dto.SubRequest) (int, error){
	var id int
	err := config.DB.QueryRow(`
        INSERT INTO subscriptions (service_name, price_month, user_uuid, start_date, end_date)
        VALUES ($1, $2, $3, $4, $5) RETURNING id_sub
				`, req.NameSub, req.PriceMonth, req.UserId, req.DateStart, req.DateEnd).Scan(&id)
	if err != nil {
			 fmt.Println("Ошибка запроса")
        return  0, err 
    }
		return  id, nil
}
