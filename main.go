package collection

import (
	"encoding/json"
	"math/rand"
	"sync"
)

// Collection is a generic collection type
type Collection[V any] struct {
	data map[string]V
	m    sync.RWMutex
}

type CollectionEntry[V any] struct {
	key   string
	value V
}

// New creates a new instance of Collection
func New[V any](size int) *Collection[V] {
	return &Collection[V]{
		data: make(map[string]V, size),
	}
}

// Set sets the value for the given key
func (c *Collection[V]) Set(key string, value V) {
	c.m.Lock()
	defer c.m.Unlock()
	c.data[key] = value
}

// Get gets the value for the given key
func (c *Collection[V]) Get(key string) V {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.data[key]
}

// Delete deletes the value for the given key
func (c *Collection[V]) Delete(key string) {
	c.m.Lock()
	defer c.m.Unlock()
	delete(c.data, key)
}

// Size returns the number of elements in the collection
func (c *Collection[V]) Size() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return len(c.data)
}

// Keys returns a slice of all the keys in the collection
func (c *Collection[V]) Keys() []string {
	c.m.RLock()
	defer c.m.RUnlock()
	keys := make([]string, 0, len(c.data))
	i := 0
	for k := range c.data {
		keys[i] = k
		i++
	}
	return keys
}

// Values returns a slice of all the values in the collection
func (c *Collection[V]) Values() []V {
	c.m.RLock()
	defer c.m.RUnlock()
	values := make([]V, len(c.data))
	i := 0
	for _, v := range c.data {
		values[i] = v
		i++
	}
	return values
}

// Each iterates over the collection and calls the callback function for each item
func (c *Collection[V]) Each(f func(key string, value V)) {
	c.m.RLock()
	defer c.m.RUnlock()
	for k, v := range c.data {
		f(k, v)
	}
}

// Clear clears the collection
func (c *Collection[V]) Clear() {
	c.m.Lock()
	defer c.m.Unlock()
	c.data = make(map[string]V)
}

// Contains returns true if the collection contains the given key
func (c *Collection[V]) Contains(key string) bool {
	c.m.RLock()
	defer c.m.RUnlock()
	_, ok := c.data[key]
	return ok
}

// Find returns 1st element that satisfies the given predicate
func (c *Collection[V]) Find(f func(key string, value V) bool) (V, bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	for k, v := range c.data {
		if f(k, v) {
			return v, true
		}
	}
	return *new(V), false
}

// Filter returns a new collection containing all the elements that satisfy the given predicate
func (c *Collection[V]) Filter(f func(key string, value V) bool) *Collection[V] {
	c.m.RLock()
	defer c.m.RUnlock()
	newC := New[V](len(c.data))
	for k, v := range c.data {
		if f(k, v) {
			newC.Set(k, v)
		}
	}
	return newC
}

// Map returns a new collection containing the results of applying the given function to each element
func (c *Collection[V]) Map(f func(key string, value V) V) *Collection[V] {
	c.m.RLock()
	defer c.m.RUnlock()
	newC := New[V](len(c.data))
	for k, v := range c.data {
		newC.Set(k, f(k, v))
	}
	return newC
}

// Reduce applies the given function to each element of the collection and returns the result
func (c *Collection[V]) Reduce(f func(acc V, value V) V, init V) V {
	c.m.RLock()
	defer c.m.RUnlock()
	acc := init
	for _, v := range c.data {
		acc = f(acc, v)
	}
	return acc
}

// Combines this collection with others into a new collection. None of the source collections are modified.
func (c *Collection[V]) Concat(others ...*Collection[V]) *Collection[V] {
	c.m.RLock()
	defer c.m.RUnlock()
	newC := New[V](len(c.data))
	for k, v := range c.data {
		newC.Set(k, v)
	}
	for _, other := range others {
		other.m.RLock()
		defer other.m.RUnlock()
		for k, v := range other.data {
			newC.Set(k, v)
		}
	}
	return newC
}

// Every checks if all items passes a test.
func (c *Collection[V]) Every(f func(key string, value V) bool) bool {
	c.m.RLock()
	defer c.m.RUnlock()
	for k, v := range c.data {
		if !f(k, v) {
			return false
		}
	}
	return true
}

// Some checks if some items passes a test.
func (c *Collection[V]) Some(f func(key string, value V) bool) bool {
	c.m.RLock()
	defer c.m.RUnlock()
	for k, v := range c.data {
		if f(k, v) {
			return true
		}
	}
	return false
}

// Entries returns a slice of CollectionEntry with the key and value of each item in the collection
func (c *Collection[V]) Entries() []CollectionEntry[V] {
	c.m.RLock()
	defer c.m.RUnlock()
	entries := make([]CollectionEntry[V], len(c.data))
	i := 0
	for k, v := range c.data {
		entries[i] = CollectionEntry[V]{k, v}
		i++
	}
	return entries
}

// Obtains random value from this collection.
func (c *Collection[V]) Random() V {
	c.m.RLock()
	defer c.m.RUnlock()
	randomKey := c.Keys()[rand.Intn(c.Size())]
	return c.data[randomKey]
}

// Sweep removes items that satisfy the provided filter function.
func (c *Collection[V]) Sweep(callback func(key string, value V) bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	for key, value := range c.data {
		if callback(key, value) {
			c.Delete(key)
		}
	}
}

// First returns the first element in the collection.
func (c *Collection[V]) First() V {
	return c.Get(c.Keys()[0])
}

// FirstN returns the first n elements n the collection
func (c *Collection[V]) FirstN(n int) *Collection[V] {
	col := New[V](n)
	keys := c.Keys()

	for i := 0; i < n; i++ {
		key := keys[i]
		value := c.Get(key)

		col.Set(key, value)
	}

	return col
}

// Last returns the last element in the collection.
func (c *Collection[V]) Last(n ...int) V {
	return c.Get(c.Keys()[c.Size()-1])
}

// LastN returns the last n elements in the collection.
func (c *Collection[V]) LastN(n int) *Collection[V] {
	col := New[V](n)
	keys := c.Keys()
	for i := len(keys) - n; i < len(keys); i++ {
		col.Set(keys[i], c.Get(keys[i]))
		n--
	}
	return col
}

/*
JSON serializes the contents of the Collection into a JSON-encoded byte slice.
It represents each element in the collection as an array containing a key (string)
and a value (of type V).
*/
func (c *Collection[V]) JSON() (*[]byte, error) {
	var jsonArray []interface{}
	for key, value := range c.data {
		jsonArray = append(jsonArray, []interface{}{key, value})
	}
	b, err := json.Marshal(jsonArray)
	return &b, err
}
