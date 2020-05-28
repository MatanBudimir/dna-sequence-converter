package main

import (
	"dna-sequence-converter/web"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal(err)
	} else if os.Getenv("USER") == "" || os.Getenv("DATABASE") == "" {
		log.Fatal("No environment variables found.")
	}

	http.HandleFunc("/", web.Index)

	http.HandleFunc("/about", web.About)

	http.HandleFunc("/convert", web.Convert)

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.ListenAndServe(":8000", nil)
}