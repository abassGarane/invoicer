package main

import (
	"fmt"
	"net/http"

	"github.com/abassGarane/learning/handlers"
	// "github.com/abassGarane/learning/pdf"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Hello, world...")
	// pdf.PdfGenerator()

	// p := pdf.NewPDFService()
	h := handlers.DefaultHandler{}

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	// css := http.ServeFile(w http.ResponseWriter, r *http.Request, "./static/index.css")
	// mux.Get("/", p.GetPDF)
	mux.Post("/", h.GenerateInvoice)
	fs := http.FileServer(http.Dir("./"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.Get("/", h.HomePage)

	http.ListenAndServe(":7071", mux)

}

func serveCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.css")
}
