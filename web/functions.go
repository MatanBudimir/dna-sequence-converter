package web

import (
	"dna-sequence-converter/helpers"
	"html/template"
	"net/http"
)

type convert struct {
	RNAValue string
	ProteinValue string
	Data *helpers.App
}

func Index(w http.ResponseWriter, r *http.Request) {
	templates, _ := template.ParseFiles("views/layout.html", "views/index.html")

	templates.Execute(w, helpers.Information(r))
}

func About(w http.ResponseWriter, r *http.Request) {
	templates, _ := template.ParseFiles("views/layout.html", "views/about.html")
	templates.Execute(w, helpers.Information(r))
}

func Convert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.FormValue("sequence") == "" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}

	templates, _ := template.ParseFiles("views/convert.html")

	converted := helpers.Convert(r.FormValue("sequence"), w, r)

	data := &convert{RNAValue: converted, Data: helpers.Information(r)}

	templates.Execute(w, data)
}