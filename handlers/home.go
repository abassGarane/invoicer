package handlers

import (
	"html/template"
	"net/http"
)

func (h *DefaultHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/partials/base.html",
		"./templates/partials/head.html",
		"./templates/partials/header.html",
		"./templates/partials/footer.html",
		"./templates/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Could not parse files", http.StatusInternalServerError)
		return
	}
	if err := ts.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Could not parse files", http.StatusInternalServerError)
		return
	}
}
