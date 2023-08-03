package helpers

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func GenerateInvoicePDF(invoiceData map[string]interface{}) []byte {
	// Create a new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Set initial position and left margin
	x, y, leftMargin := 10.0, 10.0, 10.0

	// Add the title "Spences Of Spices"
	pdf.SetXY(x, y)
	pdf.Cell(0, 10, "Spences Of Spices")
	y += 10

	// Add a line separator
	pdf.SetLineWidth(0.2)
	pdf.Line(leftMargin, y, 200-leftMargin, y)
	y += 10

	// Add invoice details to the PDF
	for key, value := range invoiceData {
		// Use SetLeftMargin to control the left margin for the text
		pdf.SetLeftMargin(leftMargin)

		// Use SetXY to set the position
		pdf.SetXY(x, y)

		// Use MultiCell to wrap text and control alignment
		switch val := value.(type) {
		case string:
			pdf.MultiCell(0, 10, key+": "+val, "0", "L", false)
		case map[string]string:
			// Handle "Delivery Address" field
			for k, v := range val {
				pdf.MultiCell(0, 10, k+": "+v, "0", "L", false)
			}
		case []map[string]interface{}:
			// Handle "Product name" field
			for _, item := range val {
				for k, v := range item {
					pdf.MultiCell(0, 10, k+": "+fmt.Sprintf("%v", v), "0", "L", false)
				}
			}
		default:
			// Handle other cases (if any)
			pdf.MultiCell(0, 10, key+": "+fmt.Sprintf("%v", value), "0", "L", false)
		}

		y += 10
	}

	// Save the PDF to a buffer
	var pdfBuf bytes.Buffer
	pdf.Output(&pdfBuf)

	return pdfBuf.Bytes()
}
