package cache

type Cache struct {
	data map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, ok := c.data[key]

	if !ok {
		return nil, false
	}

	return value, true
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func (c *Cache) Clear() {
	c.data = make(map[string]interface{})
}
