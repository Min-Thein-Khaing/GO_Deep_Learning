package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CatFactStruct struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {

	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("status code not 200", res.Status)
	}

	bodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data CatFactStruct
	if err := json.Unmarshal(bodyByte, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Println(data.Fact,data.Length)

}
