package api

import "testing"

func TestTotalPagesWithExactlyFullPages(t *testing.T) {
	if got := TotalPages(40, 20); got != 2 {
		t.Errorf("TotalPages(40, 20) = %d, want 2", got)
	}
	if got := TotalPages(100, 10); got != 10 {
		t.Errorf("TotalPages(100, 10) = %d, want 10", got)
	}
}

func TestTotalPagesHandlesEmptyAndInvalid(t *testing.T) {
	if got := TotalPages(0, 20); got != 0 {
		t.Errorf("TotalPages(0, 20) = %d, want 0", got)
	}
	if got := TotalPages(10, 0); got != 0 {
		t.Errorf("TotalPages(10, 0) = %d, want 0", got)
	}
}

func TestPageReturnsTheRequestedSlice(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	got := Page(items, 2, 3)
	if len(got) != 3 || got[0] != 4 {
		t.Errorf("Page(items, 2, 3) = %v, want [4 5 6]", got)
	}
}

func TestPagePastTheEndIsEmpty(t *testing.T) {
	items := []int{1, 2, 3}
	if got := Page(items, 9, 3); got != nil {
		t.Errorf("Page(items, 9, 3) = %v, want nil", got)
	}
}
