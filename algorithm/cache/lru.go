package main

import (
	"fmt"
)

type ListNode struct {
	value int
	key   int
	pre   *ListNode
	next  *ListNode
}

type LRU struct {
	head  *ListNode
	tail  *ListNode
	dict  map[int]*ListNode
	count int
	size  int
}

func (lru *LRU) put(key, value int) {
	if v, ok := lru.dict[key]; ok {
		v.value = value
		lru.listRemove(v)
		lru.listAddHead(v)
	} else {
		if lru.size == lru.count {
			node := lru.listRemoveLast()
			if node == nil {
				return
			} else {
				lru.size--
				delete(lru.dict, node.key)
			}
		}
		node := &ListNode{
			value: value,
			key:   key,
		}
		lru.listAddHead(node)
		lru.size++
		lru.dict[key] = node
	}
}

func (lru *LRU) get(key int) int {
	if v, ok := lru.dict[key]; ok {
		lru.listRemove(v)
		lru.listAddHead(v)
		return v.value
	}
	return -1
}

func (lru *LRU) clear(key int) {
	if v, ok := lru.dict[key]; ok {
		node := lru.listRemove(v)
		if node == nil {
			return
		} else {
			lru.size--
			delete(lru.dict, key)
		}
	}
}

//

//helper
func (lru *LRU) listAddHead(node *ListNode) {
	node.next = lru.head.next
	node.pre = lru.head
	node.next.pre = node
	node.pre.next = node
}

func (lru *LRU) listRemove(node *ListNode) *ListNode {
	if node == lru.head || node == lru.tail {
		return nil
	}
	node.pre.next = node.next
	node.next.pre = node.pre
	node.next = nil
	node.pre = nil
	return node
}

func (lru *LRU) listRemoveLast() *ListNode {
	node := lru.tail.pre
	if node == lru.head {
		return nil
	}
	return lru.listRemove(node)
}

func main() {
	head := &ListNode{
		pre: nil,
	}
	tail := &ListNode{
		next: nil,
	}
	head.next = tail
	tail.pre = head

	lru := &LRU{
		head:  head,
		tail:  tail,
		count: 2,
		size:  0,
		dict:  make(map[int]*ListNode),
	}
	lru.put(1, 1)
	lru.put(2, 2)
	lru.put(3, 3)
	fmt.Println(lru.get(1))
}
