package vat

import "fmt"

func ExampleRenderInvoice() {
	invoice := RenderInvoice(100, 0.19)
	fmt.Println(invoice)
	// Output:
	// Netto: 100.00 €
	// VAT (0.19): 19.00 €
	// Total: 119.00 €
}

func ExampleRenderInvoice_second() {
	invoice := RenderInvoice(87, 0.19)
	fmt.Println(invoice)
	// Output:
	// Netto: 87.00 €
	// VAT (0.19): 16.53 €
	// Total: 103.53 €
}
