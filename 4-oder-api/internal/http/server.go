// package httpserver модуль http сервера
package http

import (
	"context"
	"log"
	"net/http"
	"order-api/config"
	"order-api/internal/product"
	"order-api/pkg/db"
	"order-api/pkg/middleware"
)

// запуск сервера, отслеживание закрытия контекста
func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	data := db.NewDB(cfg)

	//Repo
	repo := product.NewProductRepository(data)

	//Router
	r := http.NewServeMux()

	//Handlers
	product.NewProductHandler(r, repo)

	// Logging middleware
	handlerWithLogging := middleware.LoggingMiddleware(r)

	s := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: handlerWithLogging,
	}

	go func() {
		<-ctx.Done()
		log.Println("Shutting down server ...")
		s.Shutdown(ctx)
	}()

	log.Println("Starting server on port:", cfg.ServerAddr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
