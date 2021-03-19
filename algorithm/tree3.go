package main

import "fmt"

type TreeNode3 struct {
	value string
	left  *TreeNode3
	right *TreeNode3
}

const DEFAULT = "#"

func serilize(root *TreeNode3) string {
	if root == nil {
		return DEFAULT
	}
	var result string
	result = result + root.value
	result = result + serilize(root.left) + serilize(root.right)
	return result
}

func deserilize(s []rune) (*TreeNode3, []rune) {
	a := string((s)[0])
	if a == DEFAULT {
		return nil, s
	}
	root := &TreeNode3{value: a}
	if len(s) >= 1 {
		root.left, s = deserilize(s[1:])
	}
	if len(s) >= 1 {
		root.right, s = deserilize(s[1:])
	}

	return root, s
}

func main13() {
	t4 := &TreeNode3{
		value: "4",
	}
	t3 := &TreeNode3{
		value: "3",
	}
	t2 := &TreeNode3{
		value: "2",
		left:  t4,
	}
	t1 := &TreeNode3{
		value: "1",
		left:  t2,
		right: t3,
	}

	result := serilize(t1)

	fmt.Println(result)

	resultT, a := deserilize([]rune(result))

	fmt.Println(resultT)
	fmt.Println(a)
}
