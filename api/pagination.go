// Package api contains helpers shared by the HTTP handlers.
package api

// DefaultPageSize is used when a request does not specify one.
const DefaultPageSize = 20

// TotalPages reports how many pages of the given size are needed to cover
// totalItems.
func TotalPages(totalItems, pageSize int) int {
	if pageSize <= 0 || totalItems <= 0 {
		return 0
	}
	return totalItems / pageSize
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
