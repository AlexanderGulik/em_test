package main

import (
	"em_test/src/config"
	"em_test/src/routes"
	"em_test/src/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
 if err := utils.InitLogFile(); err != nil {
        panic(err) 
    }
	fmt.Println("start server")
    
	err := godotenv.Load(".env")
	if err != nil {
    log.Println("Не удалось загрузить .env файл, используются переменные окружения")
		utils.LogError(err)
  }
	config.InitDB()
	defer config.DB.Close()
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	fmt.Println("Сервер запущен на :8080")
 	
	if err := http.ListenAndServe(":8080", mux); err != nil {
		utils.LogError(err)
        log.Fatal("Ошибка запуска сервера:", err)
    }

}
