package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	M       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeout time.Duration) *Cache {
	fmt.Println("Starting cache")
	newCache := Cache{
		Entries: make(map[string]cacheEntry),
		M:       sync.Mutex{},
	}
	go newCache.reapLoop(timeout)
	return &newCache
}

func (c *Cache) Add(key string, data []byte) {
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	c.M.Lock()
	defer c.M.Unlock()
	c.Entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.M.Lock()
	defer c.M.Unlock()
	data := c.Entries[key]
	if len(data.val) < 1 {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	for range tick.C {
		now := time.Now().UTC()
		c.reap(now, interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.M.Lock()
	for k, v := range c.Entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.Entries, k)
		}
	}
	c.M.Unlock()
}
