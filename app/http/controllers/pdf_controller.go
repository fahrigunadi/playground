package controllers

import (
	"bytes"
	"strconv"

	"codeberg.org/go-pdf/fpdf"
	"github.com/goravel/framework/contracts/http"
)

type PdfController struct {
	// Dependent services
}

func NewPdfController() *PdfController {
	return &PdfController{
		// Inject services
	}
}

func (r *PdfController) Index(ctx http.Context) http.Response {
	pageCount := ctx.Request().QueryInt("page_count", 1)

	pageCount = max(pageCount, 1)

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 14)

	for i := 1; i <= pageCount; i++ {
		pdf.AddPage()
		pdf.Cell(40, 10, "Test PDF Page "+strconv.Itoa(i))
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return ctx.Response().Json(500, http.Json{
			"message": "Internal Server Error",
		})
	}

	return ctx.Response().Data(200, "application/pdf", buf.Bytes())
}
