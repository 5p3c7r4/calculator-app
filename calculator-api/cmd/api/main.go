package main

import (
	"calculator-api/internal/handler"
	"calculator-api/internal/services"
	logger "calculator-api/pkg"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.TODO()

	customLogger := logger.New().GetJSONLogger()
	slog.SetDefault(customLogger)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page not found"})
	})

	service := services.NewCalculatorService()
	handler := handler.NewCalculatorHandler(service)

	r.POST("api/v1/calculator", handler.Calculate)

	httpAddr := ":8000"
	server := &http.Server{
		Addr:         httpAddr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// For serving frontend
	r.Static("/assets", "./frontend/assets")
	r.StaticFile("/", "./frontend/index.html")

	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	go func() {
		slog.InfoContext(ctx, "HTTP Server Listening on "+httpAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.ErrorContext(ctx, "Error on server", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "HTTP shutdown error", "error", err.Error())
	}
	slog.InfoContext(ctx, "Server Stopped Gracefully")
}
