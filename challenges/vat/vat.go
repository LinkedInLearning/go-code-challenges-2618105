package vat

import "fmt"

func RenderInvoice(price, vatRate float64) string {
	vat := price * vatRate
	total := price + vat

	return fmt.Sprintf("Netto: %.2f €\nVAT (%.2f): %.2f €\nTotal: %.2f €", price, vatRate, vat, total)
}
