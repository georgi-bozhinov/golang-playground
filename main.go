package main

import (
	"Go-Curriculum/notes"
	"log"
	"net/http"
)

func main() {
	router := notes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
