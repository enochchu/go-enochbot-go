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

	if !CheckSuccessfulHttpStatusCode(res) {
		log.Fatalf("GetResponseAsByte - Status code of %d not within 2XX range\n", res.StatusCode)
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

func CheckSuccessfulHttpStatusCode(res *http.Response) bool {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return true
	}
	return false
}
