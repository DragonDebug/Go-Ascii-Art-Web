package main

import (
	"asciiartweb/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	_ = http.ListenAndServe(":8080", nil)
}
