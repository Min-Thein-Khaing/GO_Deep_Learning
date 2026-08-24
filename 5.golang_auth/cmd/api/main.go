package main

import (
	"context"
	"go_auth/internal/app"
	httpserver "go_auth/internal/httpServer"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()

	a ,err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}
	defer func() {
		if err := a.Close(); err != nil {
			log.Printf("failed to close app: %v", err)
		}
	}()

	router := httpserver.NewRouter(a)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
		ReadHeaderTimeout:  10 * time.Second,

	}

	log.Printf("Starting server on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}