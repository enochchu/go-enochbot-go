package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetResponseAsByte(url string) []byte {
	var res, err = http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	return ResponseToJsonAsByte(res)
}

func ResponseToJsonAsByte(res *http.Response) []byte {
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return []byte(body)
}
