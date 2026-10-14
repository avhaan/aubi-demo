// Package billing turns line items into an invoice total.
//
// All money is handled in integer cents. Tax rates are expressed in basis
// points (1 bps = 0.01%), so 825 bps is 8.25%.
package billing

type LineItem struct {
	Description string
	UnitCents   int64
	Quantity    int64
	TaxRateBps  int64
}

type Invoice struct {
	Number string
	Items  []LineItem
}

// Subtotal is the pre-tax cost of the line.
func (li LineItem) Subtotal() int64 {
	return li.UnitCents * li.Quantity
}

// Tax is the tax owed on this line, in cents.
func (li LineItem) Tax() int64 {
	return li.Subtotal() * li.TaxRateBps / 10000
}

// Total is the full amount due on the invoice, in cents.
func (inv Invoice) Total() int64 {
	var total int64
	for _, item := range inv.Items {
		total += item.Subtotal() + item.Tax()
	}
	return total
}

// Subtotal is the pre-tax total across every line.
func (inv Invoice) Subtotal() int64 {
	var subtotal int64
	for _, item := range inv.Items {
		subtotal += item.Subtotal()
	}
	return subtotal
}
