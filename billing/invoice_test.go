package billing

import "testing"

func TestLineItemSubtotal(t *testing.T) {
	li := LineItem{UnitCents: 250, Quantity: 4}
	if got := li.Subtotal(); got != 1000 {
		t.Errorf("Subtotal() = %d, want 1000", got)
	}
}

func TestLineItemTax(t *testing.T) {
	li := LineItem{UnitCents: 1000, Quantity: 1, TaxRateBps: 1000} // 10%
	if got := li.Tax(); got != 100 {
		t.Errorf("Tax() = %d, want 100", got)
	}
}

func TestInvoiceSubtotal(t *testing.T) {
	inv := Invoice{Items: []LineItem{
		{UnitCents: 1000, Quantity: 2},
		{UnitCents: 500, Quantity: 1},
	}}
	if got := inv.Subtotal(); got != 2500 {
		t.Errorf("Subtotal() = %d, want 2500", got)
	}
}

func TestInvoiceTotalIncludesTax(t *testing.T) {
	inv := Invoice{Items: []LineItem{
		{UnitCents: 1000, Quantity: 1, TaxRateBps: 1000}, // 1000 + 100
		{UnitCents: 2000, Quantity: 1, TaxRateBps: 500},  // 2000 + 100
	}}
	if got := inv.Total(); got != 3200 {
		t.Errorf("Total() = %d, want 3200", got)
	}
}

func TestEmptyInvoiceTotalsZero(t *testing.T) {
	if got := (Invoice{}).Total(); got != 0 {
		t.Errorf("Total() = %d, want 0", got)
	}
}
