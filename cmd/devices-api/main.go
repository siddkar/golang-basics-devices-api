package main

import (
	"context"
	"devices-api/internal/config"
	"devices-api/internal/http/handlers/device"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Welcome to devices api")
	// load config
	cfg := config.MustLoad()

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/devices", device.CreateDevice())

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start the server")
		}
	}()

	<-done

	// Graceful shutdown logic
	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}
