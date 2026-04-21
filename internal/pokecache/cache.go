package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache{	
	output := &Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	go output.reapLoop(interval)
	return output
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) (val []byte, exist bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheEntries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.cacheEntries {
			if time.Since(value.createdAt) > interval {
				delete(c.cacheEntries, key)
			}
		} 
		c.mu.Unlock()
	}
}