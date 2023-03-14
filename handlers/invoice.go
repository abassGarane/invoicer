package handlers

import (
	"encoding/json"
	"fmt"
	// "html/template"
	"net/http"

	"github.com/abassGarane/learning/pdf"
)

// Invoice struct
type invoice struct {
	Ref      string `json:"ref"`
	Sender   string `json:"sender"`
	Reciever string `json:"reciever"`
	Amount   uint16 `json:"amount"`
	DueDate  string `json:"due_date"`
}

func (h *DefaultHandler) GenerateInvoice(w http.ResponseWriter, r *http.Request) {
	inv := &invoice{}

	if err := json.NewDecoder(r.Body).Decode(inv); err != nil {
		http.Error(w, "Could not parse Request body", http.StatusBadRequest)
	}

	fmt.Printf("\nRecieved invoice \n %v", inv)
	ret, err := pdf.NewPDFService().GeneratePDF(inv)
	if err != nil {
		http.Error(w, "Could not Generate pdf", http.StatusBadRequest)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=sample.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	// w.WriteHeader(http.StatusOK)
	w.Write(ret)
}

// func (h *DefaultHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
// 	templ := template.New("create_invoice.html")
// 	templ.
//
// }
