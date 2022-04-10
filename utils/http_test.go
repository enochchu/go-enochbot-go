package utils

import (
	"encoding/json"
	"net/http"
	"testing"
)

type TestPerson struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
}

func TestGetResponseAsByte(t *testing.T) {
	var url string = "https://swapi.dev/api/people/1"

	res := GetResponseAsByte(url)

	if res == nil {
		t.Errorf("GetResponseAsByte was not succeesful")
	}
}

func TestResponseToJsonAsByte(t *testing.T) {
	var person TestPerson
	var url string = "https://swapi.dev/api/people/1"

	res, _ := http.Get(url)

	data := ResponseToJsonAsByte(res)

	json.Unmarshal(data, &person)

	if name := person.Name; name != "Luke Skywalker" {
		t.Errorf("ResponseToJSON was not succeesful")
	}
}
