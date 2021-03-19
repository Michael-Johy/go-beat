package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

//迭代或遍历
func traverse(p *ListNode) *ListNode {
	if p == nil {
		return nil
	}
	result := p

	p = p.next

	result.next = nil

	for {
		if p != nil {
			a := p
			p = p.next
			a.next = result
			result = a
		} else {
			break
		}
	}
	return result
}

func rescru(p *ListNode) *ListNode {
	if p.next == nil {
		return p
	}
	last := rescru(p.next)
	//p = 4
	p.next.next = p
	p.next = nil
	return last
}

func reverseN(p *ListNode, n int) (*ListNode, *ListNode, *ListNode) {
	return doReverseN(p, n)
}

func doReverseN(p *ListNode, n int) (*ListNode, *ListNode, *ListNode) {
	if n == 1 || p.next == nil {
		a := p.next
		return p, nil, a
	}
	head, _, a := doReverseN(p.next, n-1)
	p.next.next = p
	p.next = a
	return head, p, a
}

func reverseByK(p *ListNode, k int) *ListNode {
	head, tail, other := doReverseN(p, k)
	for {
		fmt.Println("for")
		if tail != nil && tail.next != nil {
			tail.next, tail, other = doReverseN(other, k)
		} else {
			fmt.Println("break ")
			break
		}
	}
	return head
}

func main11() {
	ln7 := &ListNode{
		value: 7,
		next:  nil,
	}
	ln6 := &ListNode{
		value: 6,
		next:  ln7,
	}
	ln5 := &ListNode{
		value: 5,
		next:  ln6,
	}
	ln4 := &ListNode{
		value: 4,
		next:  ln5,
	}
	ln3 := &ListNode{
		value: 3,
		next:  ln4,
	}
	ln2 := &ListNode{
		value: 2,
		next:  ln3,
	}
	ln1 := &ListNode{
		value: 1,
		next:  ln2,
	}

	a := ln1
	//b := traverse(a)
	//fmt.Println(b)

	//result := rescru(a)
	//fmt.Println(result)

	head, tail, other := reverseN(a, 3)
	fmt.Println(head)
	fmt.Println(tail)
	fmt.Println(other)
	//fmt.Println(a)

	result := reverseByK(a, 3)
	fmt.Println(result)
}
