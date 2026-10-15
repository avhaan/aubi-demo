// Package api contains helpers shared by the HTTP handlers.
package api

// DefaultPageSize is used when a request does not specify one.
const DefaultPageSize = 20

// TotalPages reports how many pages of the given size are needed to cover
// totalItems.
//
// The division rounds up so that a partial final page is still counted: with
// 25 items and a page size of 20 there are 2 pages, the second holding the
// remaining 5 items. Non-positive inputs yield 0 pages.
func TotalPages(totalItems, pageSize int) int {
	if pageSize <= 0 || totalItems <= 0 {
		return 0
	}
	return (totalItems + pageSize - 1) / pageSize
}

// Page returns the slice of items on the given 1-indexed page.
func Page[T any](items []T, page, pageSize int) []T {
	if pageSize <= 0 || page <= 0 {
		return nil
	}
	start := (page - 1) * pageSize
	if start >= len(items) {
		return nil
	}
	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}
	return items[start:end]
}
