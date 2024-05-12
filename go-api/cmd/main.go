package main

import (
	"log"
	"net/http"

	"github.com/marumaro-git/study-go-api/go-api/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/study", handler.StudyPostHandler)

	err := http.ListenAndServe(":8089", mux)
	if err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
