package main

import (
	asciiartweb "asciiartweb/Pages"
	"net/http"
)

func main() {
	http.HandleFunc("/", asciiartweb.Home)
	_ = http.ListenAndServe(":8080", nil)
}
