package main

type LinkListLRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

// InitDLinkedNode 创建节点
func InitDLinkedNode(k, v int) *DLinkedNode {
	node := &DLinkedNode{
		key:   k,
		value: v,
		pre:   nil,
		next:  nil,
	}
	return node
}

// Constructor 初始化 LRUCache
func Constructor(cap int) *LinkListLRUCache {
	l := &LinkListLRUCache{
		size:     0,
		capacity: cap,
		cache:    make(map[int]*DLinkedNode),
		head:     InitDLinkedNode(0, 0),
		tail:     InitDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

// Get 获取对应 key 的值
func (lru *LinkListLRUCache) Get(key int) int {
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}
	// 移动链表
	lru.UpdateToHead(node)
	return node.value
}

// Put 将对应的 key，value 放在 lru cache 中
func (lru *LinkListLRUCache) Put(key, value int) {
	// 如果不存在，则插入新的值，否则将对应的值移动到链头
	node, ok := lru.cache[key]
	if !ok {
		newNode := InitDLinkedNode(key, value)
		if lru.size >= lru.capacity {
			// 删掉最少使用的那个
			lru.DeleteLeast()
		}
		lru.cache[key] = newNode
		lru.InsertNewHead(newNode)
	} else {
		node.value = value
		lru.UpdateToHead(node)
	}
}

// InsertNewHead 在 lru 中插入成头
func (lru *LinkListLRUCache) InsertNewHead(node *DLinkedNode) {
	tmp := lru.head.next
	lru.head.next = node
	node.pre = lru.head
	node.next = tmp
	tmp.pre = node
	lru.size++
}

// DeleteLeast 删除最少用的
func (lru *LinkListLRUCache) DeleteLeast() {
	node := lru.tail.pre
	lru.tail.pre = node.pre
	node.pre.next = node.next
	node.pre = nil
	node.next = nil
	lru.size--
	delete(lru.cache, node.key)
}

// UpdateToHead 将某个节点更新成头节点
func (lru *LinkListLRUCache) UpdateToHead(node *DLinkedNode) {
	// 先把 node 从链表上拆出来
	node.pre.next = node.next
	node.next.pre = node.pre

	// 将 head 的下一个节点缓存起来，并将 head 的 next 节点指向当前节点，当前节点指向 tmp 节点
	tmp := lru.head.next
	lru.head.next = node
	node.pre = lru.head
	node.next = tmp
	tmp.pre = node
}
