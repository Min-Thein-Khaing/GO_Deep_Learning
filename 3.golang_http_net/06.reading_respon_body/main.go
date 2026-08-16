package main

import (
	
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"

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


	bodyText := string(bodyByte)

	max := 250

	if len(bodyText) > max {
		bodyText = bodyText[:max] + "..."
	}

	fmt.Println(bodyText)


}
