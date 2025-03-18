package cache

import (
	"runtime"
	"weak"
)

type Value struct {
	Buffer [1024 * 1024]uint8 // Some big stuff.
}

type Cache struct {
	Values map[string]weak.Pointer[Value]
}

// BEGIN CACHE OMIT
func (c *Cache) Delete(key string) {
	delete(c.Values, key)
}

func (c *Cache) Put(key string, value *Value) {
	runtime.AddCleanup(value, c.Delete, key)
	c.Values[key] = weak.Make(value)
}

// END CACHE OMIT
