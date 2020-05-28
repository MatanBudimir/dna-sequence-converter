package web

import (
	"dna-sequence-converter/helpers"
	"html/template"
	"net/http"
)

type convert struct {
	RNAValue string
	ProteinValue string
	DNAValue string
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

	if (r.FormValue("type") == "DNA") {
		templates, _ := template.ParseFiles("views/convert.html")

		format := helpers.AddSpace(r.FormValue("sequence"))

		converted := helpers.ConvertRNA(format, w, r)

		protein := helpers.Protein(converted, w, r)

		data := &convert{RNAValue: converted, Data: helpers.Information(r), DNAValue: format, ProteinValue: protein}

		templates.Execute(w, data)
	} else {
		templates, _ := template.ParseFiles("views/convert.html")

		format := helpers.AddSpace(r.FormValue("sequence"))

		converted := helpers.ConvertDNA(format, w, r)

		protein := helpers.Protein(format, w, r)

		data := &convert{DNAValue: converted, Data: helpers.Information(r), RNAValue: format, ProteinValue: protein}

		templates.Execute(w, data)
	}
}