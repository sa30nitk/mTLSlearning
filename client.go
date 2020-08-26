package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("https://localhost:8443/ping")
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)

}
