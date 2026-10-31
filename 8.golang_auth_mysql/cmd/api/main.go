package main

import (
	"context"
	"fmt"
	"go_auth_mysql/internal/app"
	httpserver "go_auth_mysql/internal/httpServer"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}
	defer func() {
		if err := a.Close(); err != nil {
			log.Printf("failed to close app: %v", err)
		}
	}()

	router := httpserver.NewRouter(a)

	port := a.Config.PORT
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("Starting MySQL Auth server on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen error: %s\n", err)
	}
}

