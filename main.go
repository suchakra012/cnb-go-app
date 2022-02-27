package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to GDC Cloud Bangalore webinar on Buildpacks !!\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
