//config->db->route->run-server

package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer func() {
		if err := db.Disconnect(client); err != nil {
			fmt.Printf("mongo disconnect err : %v", err)
		}
	}()

	router := server.NewRouter()
	addr := fmt.Sprintf(":%s", cfg.ServerPORT)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server fail")
	}
	_ = database

}
