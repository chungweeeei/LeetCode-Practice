package main

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
}

func newLRU(capacity int) *LRUCache {
	return &LRUCache{}
}
