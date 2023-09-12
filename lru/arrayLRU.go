package main

type ArrayLRUCache struct {
	capacity int
	cache    []CacheNode
}

type CacheNode struct {
	key   int
	value int
}

func NewArrayLRUCache(capacity int) *ArrayLRUCache {
	return &ArrayLRUCache{
		capacity: capacity,
		cache:    make([]CacheNode, 0),
	}
}

func (lru *ArrayLRUCache) Get(key int) int {
	for index, node := range lru.cache {
		if node.key == key {
			lru.moveToFront(index)
			return node.value
		}
	}
	return -1
}

func (lru *ArrayLRUCache) Put(key int, value int) {
	for i, node := range lru.cache {
		if node.key == key {
			node.value = value
			lru.moveToFront(i)
			return
		}
	}

	// 如果缓存已满，删除最久未使用的节点
	if len(lru.cache) >= lru.capacity {
		lru.cache = lru.cache[:lru.capacity-1]
	}

	// 插入新节点到切片的开头
	lru.cache = append([]CacheNode{{key, value}}, lru.cache...)
}

func (lru *ArrayLRUCache) moveToFront(index int) {
	node := lru.cache[index]
	lru.cache = append(lru.cache[:index], lru.cache[index+1:]...)
	lru.cache = append([]CacheNode{node}, lru.cache...)
}
