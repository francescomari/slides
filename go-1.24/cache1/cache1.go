package cache

// BEGIN CACHE OMIT
type Value struct {
	Buffer [1024 * 1024]uint8 // Some big stuff.
}

type Cache struct {
	Values map[string]*Value // A lot of bigg stuff.
}

func (c *Cache) Put(key string, value *Value) {
	c.Values[key] = value
}

func (c *Cache) Get(key string) *Value {
	return c.Values[key]
}

// END CACHE OMIT
