package pokecache

import (
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

func NewCache(time.Duration) Cache {
	newCache := Cache{
		M: sync.Mutex{},
	}
	return newCache
}

func (c *Cache) delLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	defer tick.Stop()

}
