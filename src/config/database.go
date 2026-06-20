package config

import (
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"em_test/src/utils"
)

var DB *sqlx.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	var err error
	DB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Ошибка подключения", err)
		utils.LogError(err)
		return
	}

	err = DB.Ping()
	if err != nil {
		utils.LogError(err)
		fmt.Println("Ошибка пинга", err)
	}
	
	fmt.Println("Подключение к БД успешно")
}

func GetDB() *sqlx.DB {
	return DB
}
