package main

import (
	"fmt"
	"math"
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

// tree max depth
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.left)
	right := maxDepth(root.right)

	return int(math.Max(float64(left), float64(right))) + 1
}

// tree max depth => sum
func maxDepSum(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	ls, lc := maxDepSum(root.left)
	rs, rc := maxDepSum(root.right)
	if lc > rc {
		return ls + root.value, lc + 1
	} else {
		return rs + root.value, rc + 1
	}
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

func main8() {
	n10 := &TreeNode{
		value: 10,
		left:  nil,
		right: nil,
	}
	n9 := &TreeNode{
		value: 9,
		left:  nil,
		right: nil,
	}
	n8 := &TreeNode{
		value: 8,
		left:  n10,
		right: n9,
	}
	n7 := &TreeNode{
		value: 7,
		left:  nil,
		right: n8,
	}
	n6 := &TreeNode{
		value: 6,
		left:  nil,
		right: n7,
	}
	n5 := &TreeNode{
		value: 5,
		left:  nil,
		right: nil,
	}
	n4 := &TreeNode{
		value: 4,
		left:  n5,
		right: nil,
	}
	n3 := &TreeNode{
		value: 3,
		left:  nil,
		right: nil,
	}
	n2 := &TreeNode{
		value: 2,
		left:  n4,
		right: n6,
	}
	n1 := &TreeNode{
		value: 1,
		left:  n2,
		right: n3,
	}

	dfs(n1)
	bfs(n1)
	fmt.Println(maxDepth(n1))
	fmt.Println(maxDepSum(n1))

	preorders := []int{1, 2, 4, 5, 6, 3, 7}
	inorders := []int{4, 2, 6, 5, 1, 7, 3}
	result := aa(preorders, inorders)
	fmt.Println(result)
}

func aa(preorders []int, inorders []int) *TreeNode {
	if len(preorders) == 0 {
		return nil
	}
	root := &TreeNode{
		value: preorders[0],
	}
	rootI := -1
	for i, value := range inorders {
		if value == preorders[0] {
			rootI = i
		}
	}
	root.left = aa(preorders[1:rootI+1], inorders[:rootI])
	root.right = aa(preorders[rootI+1:], inorders[rootI+1:])
	return root
}
