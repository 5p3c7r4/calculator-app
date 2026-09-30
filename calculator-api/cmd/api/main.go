package main

import (
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

	// validator := validator.New()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.POST("api/v1/calculator", )

	httpAddr := ":" + os.Getenv("PORT")
	server := &http.Server{
		Addr:         httpAddr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

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