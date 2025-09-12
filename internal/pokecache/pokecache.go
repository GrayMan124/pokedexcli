package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	muc          sync.Mutex
	cacheEntries map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.muc.Lock()
	c.cacheEntries[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}
	c.muc.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.muc.Lock()
	entry, exists := c.cacheEntries[key]
	defer c.muc.Unlock()
	if exists {
		return entry.val, exists
	}
	return nil, false
}

func (c *Cache) reapLoop(dur time.Duration) {
	ticker := time.NewTicker(dur)
	defer ticker.Stop()
	for range ticker.C {
		c.muc.Lock()
		for idx, entry := range c.cacheEntries {
			if time.Since(entry.createAt) > dur {
				delete(c.cacheEntries, idx)
			}
		}
		c.muc.Unlock()
	}
	c.muc.Unlock()
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}
