package main

import (
	"dna-sequence-converter/web"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/about", web.About)
	http.HandleFunc("/convert", web.Convert)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.ListenAndServe(":8000", nil)
}