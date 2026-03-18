// package httpserver модуль http сервера
package http

import (
	"context"
	"log/slog"
	"net/http"
	"order-api/config"
	"order-api/internal/product"
	"order-api/pkg/db"
)

// запуск сервера, отслеживание закрытия контекста
func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	data := db.NewDB(cfg)

	_ = product.NewProductRepository(data.DB) //далее переменную передадим в handler

	r := http.NewServeMux()

	s := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server ...")
		s.Shutdown(ctx)
	}()

	slog.Info("Starting server ...", slog.String("addr", cfg.ServerAddr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
