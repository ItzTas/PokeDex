package cache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entry: map[string]cacheEntry{},
	}
	go c.readloop(interval)
	return c
}

type Cache struct {
	entry 	 map[string]cacheEntry
	mu 	   	 sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
} 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	r, ok := c.entry[key]
	if !ok {
		return []byte{}, false
	}
	return r.val, true
}

func (c *Cache) readloop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		t := <- ticker.C
		c.mu.Lock()
		var to_remove []string
		for key, value := range c.entry {
			if t.After(value.createdAt.Add(interval)){
				to_remove = append(to_remove, key)
			}
		}
		for _, key := range to_remove {
			delete(c.entry, key)
		}
		c.mu.Unlock()
	}
}