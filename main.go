package main

import (
	"encoding/json"
	"enochchu/go-enoch-bot/utils"
	"fmt"
)

type TestPerson struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
}

func main() {
	var person TestPerson
	var url string
	url = "https://swapi.dev/api/people/1"

	data := utils.GetResponseAsByte(url)

	json.Unmarshal(data, &person)

	fmt.Println(person.Name)
}
