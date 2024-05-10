package cache

import "sync"

type Cache struct {
	entry map[string]cacheEntry
	mu 	  sync.Mutex
}