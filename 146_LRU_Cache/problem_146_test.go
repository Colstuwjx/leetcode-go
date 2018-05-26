package main

import (
	"testing"
)

func TestLRUCacheGetPut(t *testing.T) {
	t.Log("Testing [problem 146](https://leetcode.com/problems/lru-cache)...")

	cache := Constructor(2)
	t.Log("Put (1, 1)")
	cache.Put(1, 1)
	t.Log("Current Values: ", cache.Iteral())

	t.Log("Put (2, 2)")
	cache.Put(2, 2)
	t.Log("Current Values: ", cache.Iteral())

	t.Log("Get: ", cache.Get(1)) // returns 1
	t.Log("Values: ", cache.Iteral())

	t.Log("Put (3, 3)")
	cache.Put(3, 3) // evicts key 2
	t.Log("Values: ", cache.Iteral())

	t.Log("Get: ", cache.Get(2)) // returns -1 (not found)
	t.Log("Values: ", cache.Iteral())

	t.Log("Put (4, 4)")
	cache.Put(4, 4) // evicts key 1
	t.Log("Values: ", cache.Iteral())
	t.Log("Get: ", cache.Get(1)) // returns -1 (not found)
	t.Log("Get: ", cache.Get(3)) // returns 3
	t.Log("Get: ", cache.Get(4)) // returns 4

	t.Log("Test passed!")
}
