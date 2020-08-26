package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong\n")
	})

	http.ListenAndServe(":8080", nil)
}
