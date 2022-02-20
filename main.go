package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Cloud Native folks!\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
