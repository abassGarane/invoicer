package pdf

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PDFService struct{}

func NewPDFService() *PDFService {
	return &PDFService{}
}

func (p PDFService) GeneratePDF(data any) ([]byte, error) {
	var templ *template.Template
	var err error

	// use Go's default HTML template generation tools to generate your HTML
	if templ, err = template.ParseFiles("templates/pdf.html"); err != nil {
		return nil, err
	}

	// apply the parsed HTML template data and keep the result in a Buffer
	var body bytes.Buffer
	if err = templ.Execute(&body, data); err != nil {
		return nil, err
	}

	// initalize a wkhtmltopdf generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	// enable this if the HTML file contains local references such as images, CSS, etc.
	page.EnableLocalFileAccess.Set(true)

	// add the page to your generator
	pdfg.AddPage(page)

	// manipulate page attributes as needed
	pdfg.MarginLeft.Set(30)
	pdfg.MarginRight.Set(30)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// magic
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}

func (p PDFService) GetPDF(w http.ResponseWriter, r *http.Request) {
	// .....
	data := "hello world"
	pdfBytes, err := p.GeneratePDF(&data)
	if err != nil {
		http.Error(w, "Could not load pdf", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=kittens.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}
