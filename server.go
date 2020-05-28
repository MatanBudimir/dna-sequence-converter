package main

import (
	"dna-sequence-converter/web"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	} else if os.Getenv("USER") == "" || os.Getenv("DATABASE") == "" || os.Getenv("TABLE") == "" {
		log.Fatal("No environment variables found.")
	}

	http.HandleFunc("/", web.Index)
	http.HandleFunc("/about", web.About)
	http.HandleFunc("/convert", web.Convert)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.ListenAndServe(":8000", nil)
}