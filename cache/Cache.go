package cache

import "sync"

type Cache struct {
	obj sync.Map
}

func New() *Cache {
	return &Cache{
		obj: sync.Map{},
	}
}

func (c *Cache) Set(key, value interface{}) {
	c.obj.Store(key, value)
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	m, exist := c.obj.Load(key)
	if !exist {
		return nil, false
	}
	return m, true
}
