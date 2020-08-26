package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong\n")
	})

	log.Fatal(http.ListenAndServeTLS(":8443", "server-certs/cert.pem", "server-certs/key.pem", nil))
}
