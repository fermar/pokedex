package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {

	c := Cache{}
	c.mutex.Unlock()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()

	c.cache[key] = cacheEntry{time.Now(), val}

	c.mutex.Unlock()
	return
}

func (c *Cache) Get(key string) (val []byte, hit bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	ce, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}

	return ce.val, true
}

func (c *Cache) reaploop(interval time.Duration) int {
	c.mutex.Lock()
	var count int
	count = 0
	for val, ce := range c.cache {
		if time.Now().Sub(ce.createdAt) > interval {
			delete(c.cache, val)
			count++
		}
	}
	c.mutex.Unlock()
	return count
}
