package service

import (
	"fmt"
	"em_test/src/config"
)

func (s *SubService) DeleteSubId(idSub int) (error) {
	var exit bool
	err := config.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM subscriptions WHERE id_sub = $1)`, idSub).Scan(&exit)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса на проверку существования записи")
		return err
	}
	if !exit{
		fmt.Println("Подписка не существует")
		return fmt.Errorf("Подписка %d не найдена", idSub)
	}
	_, err = config.DB.Exec(`DELETE FROM subscriptions WHERE id_sub = $1`, idSub)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса")
		return  err
	}
	return nil
	}



