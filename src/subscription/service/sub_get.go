package service
import (
		"fmt"
	"em_test/src/subscription/dto"
	"em_test/src/config"
	"time"
	"github.com/jmoiron/sqlx"
	"strings"
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
	 _, err := time.Parse("01-2006", filter.DateStart)

   if err != nil {
        return 0, fmt.Errorf("Неверный формат даты начала")
    }
		endDateReq, err := time.Parse("01-2006", filter.DateEnd)

    if err != nil {
        return 0, fmt.Errorf("Неверный формат даты окончания")
    }

	if filter.UserID != "" {
        params["user_uuid"] = filter.UserID
        conditions = append(conditions, "user_uuid = :user_uuid")
    }

    if filter.ServiceName != "" {
        params["service_name"] = filter.ServiceName
        conditions = append(conditions, "service_name = :service_name")
    }

    if filter.DateStart != "" && filter.DateEnd != "" {
        params["date_start_req"] = filter.DateStart
        params["date_end_req"] = filter.DateEnd
      	conditions = append(conditions, "start_date >= :date_start_req")
    }

		baseQuery := `SELECT id_sub, price_month, start_date, end_date FROM subscriptions `

    var subs []dto.SumType
		if len(conditions) != 0 {
      baseQuery +="WHERE " +  strings.Join(conditions, " AND ")
		}

    query, args, err := sqlx.Named(baseQuery, params)
    if err != nil {
        return 0, fmt.Errorf("ошибка подготовки запроса: %w", err)
    }
    query = config.DB.Rebind(query)
		err = config.DB.Select(&subs, query, args...)
    if err != nil {
        return 0, fmt.Errorf("ошибка получения подписок: %w", err)
    }

    totalSum := 0
   	//fmt.Println(subs) 
    for _, sub := range subs {
        startDate, err := time.Parse("01-2006", sub.DateStart)
        if err != nil {
            continue 
        }

        endDate, err := time.Parse("01-2006", sub.DateEnd)
        if err != nil {
            continue 
        }

        if startDate.After(endDateReq) {
            continue 
        }

        actualEnd := endDate
        if endDate.After(endDateReq) {
            actualEnd = endDateReq 
        }

        months := monthsBetween(startDate, actualEnd)
        
        //fmt.Println(months, sub.PriceMonth, startDate, actualEnd) 
        totalSum += sub.PriceMonth * months   
			}

    return totalSum, nil
}

func monthsBetween(start, end time.Time) int {
    years := end.Year() - start.Year()
    months := int(end.Month() - start.Month())
    return years*12 + months + 1 //считаем ли мы месяц начала подписки, если нет убрать +1
}

/*
Вариант подсчет всей суммы подписки через sql
    query := `SELECT COALESCE(SUM(
            price_month * 
            (CAST(EXTRACT(MONTH FROM AGE(
                LEAST(
                    COALESCE(TO_DATE(end_date, 'MM-YYYY'),
										TO_DATE(:date_end_req, 'MM-YYYY')),
                    TO_DATE(:date_end_req, 'MM-YYYY')
                ),
                GREATEST(
                    TO_DATE(start_date, 'MM-YYYY'),
                    TO_DATE(:date_start_req, 'MM-YYYY')
                )
							)) AS integer ) + 1)
        ),  0)
        FROM subscriptions 
        WHERE `  + 
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

*/
