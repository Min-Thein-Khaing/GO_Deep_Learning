package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/test", testHandler)
	fmt.Println("server start on Port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type TestRequest struct{
	Name string `json:"name"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "post request only",
		})
		return
	}
	defer r.Body.Close()

	var req TestRequest

	dec := json.NewDecoder(r.Body)
	

	if err := dec.Decode(&req); err !=nil {
		writeJson(w,http.StatusBadRequest,map[string]any{
			"ok": false,
			"error" : err.Error(),
		})
		return
	}

	req.Name = strings.TrimSpace(req.Name)

	if req.Name == ""{
		writeJson(w,http.StatusBadRequest,map[string]any{
			"ok": false,
			"error" : "Name is required",
		})
		return
	}

	writeJson(w,http.StatusOK,map[string]any{
		"ok":true,
		"data":req,
		"time":time.Now().UTC(),
	})

	

}
