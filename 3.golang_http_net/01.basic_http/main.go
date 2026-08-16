package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get is allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Hello from go net/http server")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running on port :8080...")

	// ListenAndServe error တက်မှသာ log.Fatal အလုပ်လုပ်ပါမည်
	log.Fatal(http.ListenAndServe(":8080", nil))
}