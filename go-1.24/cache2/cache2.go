package cache

import "weak"

type Value struct {
	Buffer [1024 * 1024]uint8 // Some big stuff.
}

type Cache struct {
	Values map[string]weak.Pointer[Value]
}

// BEGIN CACHE OMIT
func (c *Cache) Put(key string, value *Value) {
	c.Values[key] = weak.Make(value)
}

func (c *Cache) Get(key string) *Value {
	if p, ok := c.Values[key]; ok {
		return p.Value()
	}
	return nil
}

// END CACHE OMIT
