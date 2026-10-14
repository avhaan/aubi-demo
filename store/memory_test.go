package store

import "testing"

func TestSetThenGet(t *testing.T) {
	m := New(4)
	m.Set("a", "alpha")

	got, ok := m.Get("a")
	if !ok || got != "alpha" {
		t.Errorf("Get(a) = %q, %v; want alpha, true", got, ok)
	}
}

func TestGetMissingKey(t *testing.T) {
	m := New(4)
	if _, ok := m.Get("nope"); ok {
		t.Error("Get on a missing key reported ok")
	}
}

func TestDeleteRemovesTheEntry(t *testing.T) {
	m := New(4)
	m.Set("a", "alpha")
	m.Delete("a")

	if _, ok := m.Get("a"); ok {
		t.Error("entry survived Delete")
	}
	if m.Len() != 0 {
		t.Errorf("Len() = %d, want 0", m.Len())
	}
}

func TestOverwritingAKeyDoesNotGrowTheCache(t *testing.T) {
	m := New(4)
	m.Set("a", "alpha")
	m.Set("a", "alpha-2")

	if m.Len() != 1 {
		t.Errorf("Len() = %d, want 1", m.Len())
	}
	if got, _ := m.Get("a"); got != "alpha-2" {
		t.Errorf("Get(a) = %q, want alpha-2", got)
	}
}

func TestCacheStaysWithinCapacity(t *testing.T) {
	m := New(2)
	m.Set("a", "alpha")
	m.Set("b", "bravo")
	m.Set("c", "charlie")

	if m.Len() > 2 {
		t.Errorf("Len() = %d, want at most 2", m.Len())
	}
}

func TestZeroCapacityFallsBackToADefault(t *testing.T) {
	m := New(0)
	m.Set("a", "alpha")
	if _, ok := m.Get("a"); !ok {
		t.Error("value was not stored with the default capacity")
	}
}
