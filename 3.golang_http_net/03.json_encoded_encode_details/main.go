package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ok", successHandler)
	fmt.Println("server start on Port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func successHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := map[string]any{
		"ok" : true,
		"message" : "Json encode successfully",
		"dateTime":time.Now().UTC(),
	}

	json.NewEncoder(w).Encode(res)
}
