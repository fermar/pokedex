package pokecache

import (
	"time"

	"github.com/fermar/pokedex/internal/pokelog"
)

func NewCache(interval time.Duration) *Cache {

	c := Cache{}
	c.cache = make(map[string]cacheEntry)

	pokelog.Logger.Println("New cache created")
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()

	c.cache[key] = cacheEntry{time.Now(), val}

	c.mutex.Unlock()
	pokelog.Logger.Println("cache updated: entry added")
	return
}

func (c *Cache) Get(key string) (val []byte, hit bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	ce, ok := c.cache[key]
	if !ok {
		pokelog.Logger.Println("cache MISS")
		return []byte{}, false
	}
	pokelog.Logger.Println("cache HIT")
	return ce.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for {
		_, ok := <-ticker.C
		if ok {
			delcount := 0
			pokelog.Logger.Println("Analizando cache..")
			c.mutex.Lock()
			for val, ce := range c.cache {
				if time.Now().Sub(ce.createdAt) > interval {
					delete(c.cache, val)
					delcount++
				}
			}
			if delcount > 0 {
				pokelog.Logger.Printf("%v entradas de cache borradas", delcount)
			}
			c.mutex.Unlock()
		}
	}
}
