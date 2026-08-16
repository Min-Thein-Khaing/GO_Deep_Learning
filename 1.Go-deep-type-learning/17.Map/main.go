package main

import (
	"fmt"
)

func main() {
	//map[keyType] valueType
	//maps are unordered collection of key value pairs

	countryCode := map[string]int{
		"Pakistan": 343,
		"India":    91,
		"USA":      1,
	}

	fmt.Println(countryCode)
	fmt.Println(countryCode["Pakistan"])

	var language map[string]string

	language = make(map[string]string)
	fmt.Println(len(language))

	language["JS"] = "Javscript"
	language["PY"] = "Python"
	language["GO"] = "Golang"

	fmt.Println(language)

	fmt.Println(len(language))

	users := map[string]int{
		"ali":     20,
		"Hania":   21,
		"Ahmed":   22,
	}
	delete(users, "ali")
	fmt.Println(users)

}
