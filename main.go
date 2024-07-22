package main

import (
	"fmt"
	"net/http"

	asciitoweb "github.com/ahatdemirezen/docker/AsciiToWeb"
)

func main() {
	http.Handle("/templates/style.css", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.HandleFunc("/", asciitoweb.AsciiToWeb)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Something went wrong!")
		return
	}
}
