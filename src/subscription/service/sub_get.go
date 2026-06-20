package service
import (
		"fmt"
	"em_test/src/subscription/dto"
	"em_test/src/config"
	"strings"

	"github.com/jmoiron/sqlx"
)

func (s * SubService) GetSubId(id int) (dto.SubResponse, error) {
	var res dto.SubResponse
	 err :=	 config.DB.Get(&res, `
    SELECT id_sub, service_name, price_month, user_uuid, start_date, end_date 
    FROM subscriptions 
    WHERE id_sub = $1
`, id)

	if err != nil {
		fmt.Println("Ошибка запроса")
		return res, err
	}
	return res, nil
}

func (s *SubService) SelectSubAll() ([]dto.SubResponse, error) {
	var res []dto.SubResponse
	err := config.DB.Select(&res, ` SELECT id_sub, service_name, price_month, user_uuid, start_date, end_date 
    FROM subscriptions 
		`)
	if err != nil {
		fmt.Println("Ошибка запроса")
		return res, err
	}
	return res, err 
}

func (s *SubService) TotalGetSum(filter dto.TotalSumRequest) (int, error) {
    params := map[string]interface{}{}
    var conditions []string

    if filter.UserID != "" {
        params["user_uuid"] = filter.UserID
        conditions = append(conditions, "user_uuid = :user_uuid")
    }

    if filter.ServiceName != "" {
        params["service_name"] = filter.ServiceName
        conditions = append(conditions, "service_name = :service_name")
    }

    if filter.DateStart != "" && filter.DateEnd != "" {
        params["start_date"] = filter.DateStart
        params["end_date"] = filter.DateEnd
        conditions = append(conditions, "TO_DATE(start_date, 'MM-YYYY') <= TO_DATE(:end_date, 'MM-YYYY')")
        conditions = append(conditions, "(TO_DATE(end_date, 'MM-YYYY') >= TO_DATE(:start_date, 'MM-YYYY') OR end_date IS NULL)")
    }

    if len(conditions) == 0 {
        return 0, fmt.Errorf("не передано ни одного фильтра")
    }

    query := "SELECT COALESCE(SUM(price_month), 0) FROM subscriptions WHERE " + 
             strings.Join(conditions, " AND ")

    query, args, err := sqlx.Named(query, params)
    if err != nil {
        return 0, fmt.Errorf("ошибка подготовки запроса: %w", err)
    }
    query = config.DB.Rebind(query)

    var total int
    err = config.DB.Get(&total, query, args...)
    if err != nil {
        return 0, fmt.Errorf("ошибка подсчёта суммы: %w", err)
    }

    return total, nil
}
