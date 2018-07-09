package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFUCache(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1))  // returns 1
	cache.Put(3, 3)                   // evicts key 2
	assert.Equal(t, -1, cache.Get(2)) // returns -1 (not found)
	assert.Equal(t, 3, cache.Get(3))  // returns 3.
	cache.Put(4, 4)                   // evicts key 1.
	assert.Equal(t, -1, cache.Get(1)) // returns -1 (not found)
	assert.Equal(t, 3, cache.Get(3))  // returns 3
	assert.Equal(t, 4, cache.Get(4))  // returns 4
}
