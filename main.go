package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	address := os.Getenv("HTTP_ADDR")
	if address == "" {
		address = ":8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request to " + r.URL.Path + " from " + r.RemoteAddr)

		w.WriteHeader(200)
		if _, err := w.Write([]byte("200 OK\n")); err != nil {
			log.Printf("Error writing response: %s\n", err)
		}
	})

	log.Printf("Listening on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
