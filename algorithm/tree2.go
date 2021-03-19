package main

import (
	"fmt"
)

type TreeNode2 struct {
	value int
	left  *TreeNode2
	right *TreeNode2
}

func build(arr []int) *TreeNode2 {
	if len(arr) == 0 {
		return nil
	}
	mi := 0
	max := arr[mi]
	for i, item := range arr {
		if item > max {
			mi = i
			max = item
		}
	}
	root := &TreeNode2{value: max}
	root.left = build(arr[0:mi])

	if mi < len(arr)-1 {
		root.right = build(arr[mi+1:])
	}

	return root
}

func main1() {
	arr := []int{3, 2, 1, 6, 0, 5}

	arr1 := []int{3}

	fmt.Println(arr1[0:])
	result := build(arr)
	fmt.Println(result)
}
