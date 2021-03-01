package lru

import (
	"errors"

	"github.com/hashicorp/golang-lru/simplelru"
	"github.com/laotoutou/tslru/tslru"
)

// TSCache is a thread-safe fixed size LRU cache.
type TSCache struct {
	// Fragmentation can reduce lock contention, but the hash function affects efficiency
	lru simplelru.LRUCache
}

// NewTSCache creates an LRU of the given size.
func NewTSCache(size int) (c *TSCache, err error) {
	if size < 1 {
		return nil, errors.New("size cannot be less than 1")
	}

	// create a cache with default settings
	lru, err := tslru.NewLRU(size)
	if err != nil {
		return nil, err
	}
	return &TSCache{lru: lru}, nil
}

// Purge is used to completely clear the cache.
func (c *TSCache) Purge() {
	c.lru.Purge()
}

// Add adds a value to the cache. Returns true if an eviction occurred.
func (c *TSCache) Add(key string, value interface{}) (evicted bool) {
	return c.lru.Add(key, value)
}

// Get looks up a key's value from the cache.
func (c *TSCache) Get(key string) (value interface{}, ok bool) {
	return c.lru.Get(key)
}

// Contains checks if a key is in the cache, without updating the
// recent-ness or deleting it for being stale.
func (c *TSCache) Contains(key string) bool {
	return c.lru.Contains(key)
}

// Peek returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (c *TSCache) Peek(key string) (value interface{}, ok bool) {
	return c.lru.Peek(key)
}

// ContainsOrAdd checks if a key is in the cache without updating the
// recent-ness or deleting it for being stale, and if not, adds the value.
// Returns whether found and whether an eviction occurred.
func (c *TSCache) ContainsOrAdd(key string, value interface{}) (ok, evicted bool) {
	if c.lru.Contains(key) {
		return true, false
	}
	return false, c.lru.Add(key, value)
}

// PeekOrAdd checks if a key is in the cache without updating the
// recent-ness or deleting it for being stale, and if not, adds the value.
// Returns whether found and whether an eviction occurred.
func (c *TSCache) PeekOrAdd(key string, value interface{}) (previous interface{}, ok, evicted bool) {
	previous, ok = c.lru.Peek(key)
	if ok {
		return previous, true, false
	}
	return nil, false, c.lru.Add(key, value)
}

// Remove removes the provided key from the cache.
func (c *TSCache) Remove(key string) (present bool) {
	return c.lru.Remove(key)
}

// Resize changes the cache size.
func (c *TSCache) Resize(size int) (evicted int) {
	return c.lru.Resize(size)
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *TSCache) Keys() (ret []string) {
	for _, k := range c.lru.Keys() {
		ret = append(ret, k.(string))
	}
	return
}

// Len returns the number of items in the cache.
func (c *TSCache) Len() (ret int) {
	return c.lru.Len()
}
