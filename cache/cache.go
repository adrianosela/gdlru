package cache

import (
	"container/list"
	"errors"
	"log"
)

// Cache is an in-memory LRU, write-back cache
type Cache struct {
	size      int
	seen      map[interface{}]*list.Element // required for constant-time access
	stack     *list.List                    // required for order keeping
	evictFunc func(k interface{}, v interface{}) error
}

// kv represents a key-value pair to
// be used as an element in the stack
type kv struct {
	k interface{}
	v interface{}
}

// NewCache returns an initialized Cache of a given size
func NewCache(size int, evictFunc func(k interface{}, v interface{}) error) (*Cache, error) {
	if evictFunc == nil {
		return nil, errors.New("evictFunc cannot be nil")
	}
	if size <= 0 || size > 1000 {
		return nil, errors.New("size must be an integer between 1 and 1000")
	}

	return &Cache{
		size:      size,
		stack:     list.New(),
		seen:      make(map[interface{}]*list.Element),
		evictFunc: evictFunc,
	}, nil
}

// Get fetches the item associated with a given key from the cache
func (c *Cache) Get(k interface{}) (v interface{}, ok bool) {
	if elem, seen := c.seen[k]; seen {
		// cache hit
		c.stack.MoveToFront(elem)
		if keyValue := elem.Value.(*kv); keyValue != nil {
			return keyValue.v, true
		}
	}
	return nil, false
}

// Put puts an item in the cache
func (c *Cache) Put(k, v interface{}) {
	if elem, seen := c.seen[k]; seen {
		// cache hit
		c.stack.MoveToFront(elem)
		elem.Value.(*kv).v = v
	} else {
		// cache miss
		newElem := c.stack.PushFront(&kv{k, v})
		c.seen[k] = newElem

		if c.stack.Len() > c.size {
			// evict
			if lru := c.stack.Back(); lru != nil {
				c.stack.Remove(lru)
				lruElem := lru.Value.(*kv)
				delete(c.seen, lruElem.k)
				c.evictFunc(lruElem.k, lruElem.v)
			}
		}
	}
}

// Commit evicts every item from the cache
func (c *Cache) Commit() {
	for key, value := range c.seen {
		if err := c.evictFunc(key, value.Value.(*kv).v); err != nil {
			log.Printf("failure evicting item from cache: %s", err)
		}
		delete(c.seen, key)
	}
	c.stack = list.New()
}
