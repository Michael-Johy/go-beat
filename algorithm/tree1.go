package main

import "fmt"

type MTreeNode struct {
	value int
	left  *MTreeNode
	right *MTreeNode
	next  *MTreeNode
}

func connect(root *MTreeNode) {
	tryConnect(root.left, root.right)
}

func tryConnect(left *MTreeNode, right *MTreeNode) {
	if left == nil || right == nil {
		return
	}
	left.next = right
	tryConnect(left.left, left.right)
	tryConnect(right.left, right.right)
	tryConnect(left.right, right.left)
}

func flatten(root *MTreeNode) {
	if root == nil {
		return
	}
	flatten(root.left)
	flatten(root.right)

	left := root.left
	right := root.right

	root.left = nil
	if left != nil {
		root.right = left
		for {
			if left.right == nil {
				left.right = right
				break
			} else {
				left = left.right
			}
		}
	}
}

func main14() {
	t7 := &MTreeNode{
		value: 7,
	}
	t6 := &MTreeNode{
		value: 6,
	}
	t5 := &MTreeNode{
		value: 5,
	}
	t4 := &MTreeNode{
		value: 4,
	}
	t3 := &MTreeNode{
		value: 3,
		left:  t6,
		right: t7,
	}
	t2 := &MTreeNode{
		value: 2,
		left:  t4,
		right: t5,
	}
	t1 := &MTreeNode{
		value: 1,
		left:  t2,
		right: t3,
	}

	t := t1
	//connect(t)
	//fmt.Println(t)
	flatten(t)
	fmt.Println("aa")
}
