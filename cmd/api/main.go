package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/UberionAI/coord-gin-rest/internal/config"
	"github.com/UberionAI/coord-gin-rest/internal/db"
	"github.com/UberionAI/coord-gin-rest/internal/handler"
	"github.com/UberionAI/coord-gin-rest/internal/logger"
	"github.com/UberionAI/coord-gin-rest/internal/middleware"
	"github.com/UberionAI/coord-gin-rest/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Инициализируем логгер
	logger.Init(cfg.LogLevel)
	log.Info().Msg("Starting coord-gin-rest service")

	// Устанавливаем режим Gin
	gin.SetMode(cfg.GinMode)

	// Инициализируем слои
	dbClient := db.NewClient(cfg.ESHost)
	svc := service.New(dbClient)
	h := handler.New(svc)

	// Создаём роутер
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LoggerMiddleware())

	// Регистрируем эндпоинты
	router.GET("/health", h.Health)
	router.GET("/v1/api", h.GetData)

	// Создаём HTTP сервер
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}

	// Запускаем сервер в горутине
	go func() {
		log.Info().Str("port", cfg.ServerPort).Msg("Server starting")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited gracefully")
}
