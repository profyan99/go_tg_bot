package cache

import (
	"go-tg-bot/helpers"
	"sync"
)

type Cacher interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Cache struct {
	mutex     sync.RWMutex
	valuesMap map[string]string
}

func NewCache() Cacher {
	return &Cache{
		valuesMap: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.valuesMap[key] = value

	return nil
}

func (c *Cache) Get(key string) (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if value, ok := c.valuesMap[key]; ok {
		return value, nil
	}

	return "", helpers.ErrNotFound
}

func (c *Cache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.valuesMap, key)

	return nil
}
