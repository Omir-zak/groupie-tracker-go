package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}

	body,err := ioutil.ReadAll(resp.Body)
	err != nil {
		log.Fatalln(err)
	}

	log.print
}