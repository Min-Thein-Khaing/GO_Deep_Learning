package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type CatFactStruct struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (*CatFactStruct, error) {

	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status code not 200")
	}

	bodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data CatFactStruct
	if err := json.Unmarshal(bodyByte, &data); err != nil {
		return nil, err
	}

	return &data, nil

}

func main() {

	http.HandleFunc("/external", externalHandler)
	fmt.Println("server start on Port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func externalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "method not allowed",
		})
		return
	}

	data, err := fetchCatFact()
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"ok" : true,
		"time" : time.Now().UTC(),
		"external":map[string]any{
			"source":"https://catfact.ninja/fact",
			"message":"cat fact loaded",
			"fact":data.Fact,
			"length":data.Length,
		},
	})

}
