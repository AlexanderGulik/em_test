package service

import (
	"fmt"
	"em_test/src/config"
)

func (s *SubService) DeleteSubId(idSub int) (error) {
	_, err := config.DB.Exec(`DELETE FROM subscriptions WHERE id_sub = $1`, idSub)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса")
		return  err
	}
	return nil
}


