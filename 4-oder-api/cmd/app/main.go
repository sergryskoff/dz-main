package main

import (
	"context"
	"log"
	"order-api/internal/http"
	"os"
	"os/signal"
)

// Выполняем программу, в случае сбоя - выходим по ошибке и логируем
func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

// загружаем http сервер в контексте, ждем события, ошибки поднимаем вверх
func realMain() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := http.Run(ctx); err != nil {
		return err
	}
	return nil

}
