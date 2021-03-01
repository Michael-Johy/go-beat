package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

/**
		1
	2		5
3		4
*/

func doDfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.value)
	if root.left != nil {
		doDfs(root.left, result)
	}
	if root.right != nil {
		doDfs(root.right, result)
	}
}

// depth first search
func dfs(root *TreeNode) {
	var result []int
	doDfs(root, &result)
	fmt.Println(result)
}

func bfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	pre := []TreeNode{*root}
	var result []int

	doBfs(&pre, &result)

	return result
}

/**
  slices used as queue
  time complexity  : O(n)
  space complexity : O(n)

*/
func doBfs(result *[]TreeNode, resultInt *[]int) {
	if len(*result) == 0 {
		return
	}

	node := (*result)[0]
	*resultInt = append(*resultInt, node.value)

	if node.left != nil {
		*result = append(*result, *node.left)
	}
	if node.right != nil {
		*result = append(*result, *node.right)
	}
	*result = (*result)[1:]

	doBfs(result, resultInt)
}

func main() {
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

	dfs(n1)
	bfs(n1)
}
