package main

// 单链表实现
type DoubleLinkedListHashMap struct {
	m   []*DoubleLinkedListNode //hash表的桶
	cap int
	len int
}

type DoubleLinkedListNode struct { //键值对
	key   string
	value interface{}
	pre   *DoubleLinkedListNode
	next  *DoubleLinkedListNode
}

// 初始化一个hashmap
func NewDoubleLinkedListHashMap(cap int) *DoubleLinkedListHashMap {
	if cap < 16 {
		cap = 16
	}
	hashMap := new(DoubleLinkedListHashMap)
	hashMap.m = make([]*DoubleLinkedListNode, cap, cap)
	hashMap.cap = cap
	return hashMap
}

func (h *DoubleLinkedListHashMap) Index(key string) int {
	return BKDRHash(key, h.cap) //计算键的哈希值
}

// 写入一个键值对
func (h *DoubleLinkedListHashMap) Put(key string, value interface{}) {
	index := h.Index(key)
	element := h.m[index]
	if element == nil { //该下标没有值，直接写入
		pre, next := &DoubleLinkedListNode{}, &DoubleLinkedListNode{}
		node := &DoubleLinkedListNode{
			key:   key,
			value: value,
			pre:   pre,
			next:  next,
		}
		pre.next = node
		next.pre = node
		h.m[index] = pre
	} else { //有值，找到最后一个节点
		for element.next != nil {
			if element.key == key { //如果是相同的键就进行修改操作
				element.value = value
				return
			}
			element = element.next
		}
		n := &DoubleLinkedListNode{
			key:   key,
			value: value,
			pre:   element,
			next:  nil,
		}
		element.next = n
	}
	h.len++
	if (float64(h.len) / float64(h.cap)) > expandFactor { //需要扩容，把切片容量变为2倍
		newH := NewDoubleLinkedListHashMap(2 * h.cap)
		for _, pairs := range h.m {
			for pairs != nil {
				newH.Put(pairs.key, pairs.value)
			}
		}
		h.cap = newH.cap
		h.m = newH.m
	}
}

// 获取值
func (h *DoubleLinkedListHashMap) Get(key string) interface{} {
	index := h.Index(key)
	pair := h.m[index]
	for pair != nil {
		if pair.key == key {
			return pair.value
		}
		pair = pair.next
	}
	return nil
}

func (h *DoubleLinkedListHashMap) Remove(key string) {
	index := h.Index(key)
	pair := h.m[index]
	if pair == nil {
		return
	}
	for e := pair.pre.next; e != nil; e = e.next {
		if e.key == key {
			e.pre.next = e.next
			e.next.pre = e.pre
		}
	}
}
