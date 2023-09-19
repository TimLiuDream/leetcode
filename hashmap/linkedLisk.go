package main

// 单链表实现
type LinkedListHashMap struct {
	m   []*LinkedListKeyPairs //hash表的桶
	cap int
	len int
}

type LinkedListKeyPairs struct { //键值对
	key   string
	value interface{}
	next  *LinkedListKeyPairs
}

// 初始化一个hashmap
func NewLinkedListHashMap(cap int) *LinkedListHashMap {
	if cap < 16 {
		cap = 16
	}
	hashMap := new(LinkedListHashMap)
	hashMap.m = make([]*LinkedListKeyPairs, cap, cap)
	hashMap.cap = cap
	return hashMap
}

func (h *LinkedListHashMap) Index(key string) int {
	return BKDRHash(key, h.cap) //计算键的哈希值
}

// 写入一个键值对
func (h *LinkedListHashMap) Put(key string, value interface{}) {
	index := h.Index(key)
	element := h.m[index]
	if element == nil { //该下标没有值，直接写入
		h.m[index] = &LinkedListKeyPairs{key, value, nil}
	} else { //有值，找到最后一个节点
		for element.next != nil {
			if element.key == key { //如果是相同的键就进行修改操作
				element.value = value
				return
			}
			element = element.next
		}
		element.next = &LinkedListKeyPairs{key, value, nil}
	}
	h.len++
	if (float64(h.len) / float64(h.cap)) > expandFactor { //需要扩容，把切片容量变为2倍
		newH := NewLinkedListHashMap(2 * h.cap)
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
func (h *LinkedListHashMap) Get(key string) interface{} {
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

func (h *LinkedListHashMap) Remove(key string) {
	index := h.Index(key)
	element := h.m[index]
	dummy := &LinkedListKeyPairs{next: element}
	prev := dummy

	for element != nil {
		if element.key == key {
			prev.next = element.next
			h.len--
			return
		}
		prev = element
		element = element.next
	}

	h.m[index] = dummy.next
}
