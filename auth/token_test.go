package auth

import (
	"testing"
	"time"
)

func newTestCache(t *testing.T) (*TokenCache, *time.Time, *int) {
	t.Helper()
	clock := time.Date(2026, 1, 1, 9, 0, 0, 0, time.UTC)
	issued := 0
	c := NewTokenCache()
	c.now = func() time.Time { return clock }
	c.issue = func(userID string) (string, error) {
		issued++
		return "token_" + userID, nil
	}
	return c, &clock, &issued
}

func TestGetOrRefreshIssuesATokenForANewUser(t *testing.T) {
	c, _, issued := newTestCache(t)

	token, err := c.GetOrRefresh("alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "token_alice" {
		t.Errorf("got token %q, want %q", token, "token_alice")
	}
	if *issued != 1 {
		t.Errorf("issued %d tokens, want 1", *issued)
	}
}

func TestGetOrRefreshReusesAFreshToken(t *testing.T) {
	c, _, issued := newTestCache(t)

	first, _ := c.GetOrRefresh("alice")
	second, _ := c.GetOrRefresh("alice")

	if first != second {
		t.Errorf("token changed between calls: %q then %q", first, second)
	}
	if *issued != 1 {
		t.Errorf("issued %d tokens for one user inside the TTL, want 1", *issued)
	}
}

func TestGetOrRefreshRejectsAnEmptyUser(t *testing.T) {
	c, _, _ := newTestCache(t)

	if _, err := c.GetOrRefresh(""); err != ErrUnknownUser {
		t.Errorf("got error %v, want ErrUnknownUser", err)
	}
}

func TestInvalidateForcesAFreshToken(t *testing.T) {
	c, _, issued := newTestCache(t)

	c.GetOrRefresh("alice")
	c.Invalidate("alice")
	c.GetOrRefresh("alice")

	if *issued != 2 {
		t.Errorf("issued %d tokens, want 2 after invalidation", *issued)
	}
}

// A token is only valid for TokenTTL. Once it has expired the cache must mint a
// new one; serving the stale value is what produces 401s from the API gateway
// after a token has aged out.
func TestGetOrRefreshReplacesAnExpiredToken(t *testing.T) {
	c, clock, issued := newTestCache(t)

	c.GetOrRefresh("alice")
	*clock = clock.Add(TokenTTL + time.Minute)
	c.GetOrRefresh("alice")

	if *issued != 2 {
		t.Errorf("issued %d tokens, want 2 — the expired token should have been replaced", *issued)
	}
}
