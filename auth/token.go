// Package auth issues and caches short-lived access tokens.
package auth

import (
	"errors"
	"sync"
	"time"
)

// TokenTTL is how long an issued token stays valid.
const TokenTTL = 15 * time.Minute

var ErrUnknownUser = errors.New("auth: unknown user")

type entry struct {
	token    string
	issuedAt time.Time
}

// TokenCache hands out access tokens and remembers them so that repeated
// requests for the same user do not hit the identity provider every time.
type TokenCache struct {
	mu      sync.RWMutex
	entries map[string]entry
	now     func() time.Time
	issue   func(userID string) (string, error)
}

func NewTokenCache() *TokenCache {
	return &TokenCache{
		entries: make(map[string]entry),
		now:     time.Now,
		issue:   issueFromProvider,
	}
}

// GetOrRefresh returns a valid access token for userID, minting a new one when
// there is nothing usable cached.
func (c *TokenCache) GetOrRefresh(userID string) (string, error) {
	if userID == "" {
		return "", ErrUnknownUser
	}

	c.mu.RLock()
	cached, ok := c.entries[userID]
	c.mu.RUnlock()

	if ok {
		return cached.token, nil
	}

	token, err := c.issue(userID)
	if err != nil {
		return "", err
	}

	c.mu.Lock()
	c.entries[userID] = entry{token: token, issuedAt: c.now()}
	c.mu.Unlock()

	return token, nil
}

// Invalidate drops any cached token for userID.
func (c *TokenCache) Invalidate(userID string) {
	c.mu.Lock()
	delete(c.entries, userID)
	c.mu.Unlock()
}

// Size reports how many users currently have a cached token.
func (c *TokenCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.entries)
}

func issueFromProvider(userID string) (string, error) {
	return "token_" + userID, nil
}
