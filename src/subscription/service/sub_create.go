package service
import (
	"fmt"
	"em_test/src/subscription/dto"
	"em_test/src/config"
)
func (s *SubService) CreateSub(req dto.SubRequest) (error){
	_, err := config.DB.Exec(`
        INSERT INTO subscriptions (service_name, price_month, user_uuid, start_date, end_date)
        VALUES ($1, $2, $3, $4, $5)
				`, req.NameSub, req.PriceMonth, req.UserId, req.DateStart, req.DateEnd)
	if err != nil {
			 fmt.Println("Ошибка запроса")
        return  err 
    }
		return  nil
}
