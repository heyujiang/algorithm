package data_structure_design

import (
	"fmt"
	"testing"
)

func TestNewLRU(t *testing.T) {
	lru := NewLRU(4)

	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	printCache(lru)

	fmt.Println(lru.Get(5))
	printCache(lru)

	fmt.Println(lru.Get(1))
	printCache(lru)

	lru.Put(5, 5)
	printCache(lru)

	lru.Put(6, 6)
	printCache(lru)

	fmt.Println(lru.Get(1))
	printCache(lru)

	fmt.Println(lru.Get(2))
	printCache(lru)

	fmt.Println(lru.Get(3))
	printCache(lru)

	fmt.Println(lru.Get(4))
	printCache(lru)
}

func printCache(lru *LRUCache) {
	head := lru.cache.head.next
	for head != lru.cache.tail {
		fmt.Println(head)
		head = head.next
	}
	fmt.Println(lru.km)
	fmt.Println("+++++++++++++++++++++")
}

func TestAA(t *testing.T) {
	a := []int{1, 2, 3, 4}
	for _, v := range a {
		fmt.Println(v)
		v = 4
	}
	fmt.Println(a)

	type node struct {
		val int
	}

	b := []*node{&node{val: 1}, &node{val: 2}}
	for _, v := range b {
		fmt.Println(v)
		v.val = 5
	}
	for _, v := range b {
		fmt.Println(v)
	}

	for i := range a {
		a[i] = 5
	}
	fmt.Println(a)
}
