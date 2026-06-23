package main

// @title           Subscription API
// @version         1.0
// @description     API для управления подписками
// @host            localhost:8080
// @BasePath        /

import (
	"em_test/src/config"
	"em_test/src/routes"
	"em_test/src/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"context"

	"github.com/joho/godotenv"
)

func main() {
	if err := utils.InitLogFile(); err != nil {
		panic(err)
	}

	fmt.Println("start server")

	if err := godotenv.Load(".env"); err != nil {
		log.Println("Не удалось загрузить .env файл")
		utils.LogError(err)
	}

	config.InitDB()
	defer config.DB.Close()

	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	fs := http.FileServer(http.Dir("docs"))
	mux.Handle("GET /swagger/", http.StripPrefix("/swagger", fs))

	server := &http.Server{Addr: ":8080", Handler: mux}

	go func() {

	fmt.Println("Сервер запущен на :8080")
	fmt.Println("Swagger UI: http://localhost:8080/swagger/index.html")

	if err := server.ListenAndServe(); err != nil {
		utils.LogError(err)
		log.Fatal("Ошибка запуска сервера:", err)
	}
}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)


	<-stop
	fmt.Println("\nОстановка сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
		if err := server.Shutdown(ctx); err != nil {
		utils.LogError(err)
		fmt.Println("Ошибка при остановке:", err)
	}

	fmt.Println("Сервер остановлен")

}
