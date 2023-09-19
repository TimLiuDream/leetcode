package main

import "container/list"

var base = 1000

type pair struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}

func (m *MyHashMap) Put(key, value int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		p := e.Value.(pair)
		if p.key == key {
			e.Value = &pair{
				key:   key,
				value: value,
			}
			return
		}
	}
	m.data[h].PushBack(&pair{key: key, value: value})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		p := e.Value.(pair)
		if p.key == key {
			return p.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		p := e.Value.(pair)
		if p.key == key {
			m.data[h].Remove(e)
		}
	}
}
