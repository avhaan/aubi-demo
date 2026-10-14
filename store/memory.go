// Package store is a small bounded in-memory cache used by the API layer.
package store

import "sync"

// Entry is a stored value plus the insertion order used for eviction.
type Entry struct {
	Value string
	Seq   int64
}

// Memory is a fixed-capacity cache. When it is full, inserting a new key
// evicts the least recently inserted entry.
type Memory struct {
	mu       sync.Mutex
	entries  map[string]Entry
	capacity int
	seq      int64
}

func New(capacity int) *Memory {
	if capacity <= 0 {
		capacity = 128
	}
	return &Memory{entries: make(map[string]Entry), capacity: capacity}
}

// Set stores value under key, evicting if the cache is at capacity.
func (m *Memory) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.entries[key]; !exists && len(m.entries) >= m.capacity {
		m.evictLocked()
	}
	m.seq++
	m.entries[key] = Entry{Value: value, Seq: m.seq}
}

// Get returns the value stored under key.
func (m *Memory) Get(key string) (string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	entry, ok := m.entries[key]
	return entry.Value, ok
}

// Delete removes key from the cache.
func (m *Memory) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.entries, key)
}

// Len reports how many entries are held.
func (m *Memory) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.entries)
}

// evictLocked drops one entry to make room. The caller must hold m.mu.
func (m *Memory) evictLocked() {
	var chosen string
	var chosenSeq int64
	for key, entry := range m.entries {
		if chosen == "" || entry.Seq > chosenSeq {
			chosen, chosenSeq = key, entry.Seq
		}
	}
	if chosen != "" {
		delete(m.entries, chosen)
	}
}
