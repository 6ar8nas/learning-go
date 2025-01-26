package utils

import "sync"

type Cache[T comparable, Y any] struct {
	store map[T]Y
	mux   sync.RWMutex
}

func NewCache[T comparable, Y any]() *Cache[T, Y] {
	return &Cache[T, Y]{make(map[T]Y), sync.RWMutex{}}
}
func (c *Cache[T, Y]) Set(id T, value Y) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.store[id] = value
}
func (c *Cache[T, Y]) Get(id T) (val Y, ok bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	val, ok = c.store[id]
	return
}
