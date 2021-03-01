package main

import "fmt"

type Stack []TreeNode

func (s *Stack) Push(i TreeNode) {
	*s = append(*s, i)
}

func (s *Stack) Pop() TreeNode {
	return (*s)[len(*s)-1]
}

func main3() {
	n5 := &TreeNode{
		value: 5,
		left:  nil,
		right: nil,
	}
	n4 := &TreeNode{
		value: 4,
		left:  nil,
		right: nil,
	}
	n3 := &TreeNode{
		value: 3,
		left:  nil,
		right: nil,
	}
	n2 := &TreeNode{
		value: 2,
		left:  n3,
		right: n4,
	}
	n1 := &TreeNode{
		value: 1,
		left:  n2,
		right: n5,
	}

	stack := &Stack{}
	stack.Push(*n1)
	stack.Push(*n2)
	stack.Push(*n3)

	node := stack.Pop()
	fmt.Println(node.right)

}
