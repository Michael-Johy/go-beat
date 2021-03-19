package main

import "fmt"

//或者是一棵空的二叉树，或者是具有下列性质的二叉树：
//1、若它的左子树不空，则左子树上所有结点的值均小于根结点的值；
//2、若它的右子树不空，则右子树上所有结点的值均大于根结点的值；
//3、它的左右子树也都是二叉排序树。
//4、中序遍历可以得到一个递增序列

type TreeNode9 struct {
	value int
	left  *TreeNode9
	right *TreeNode9
}

var a = 0
var b = -1

func findK(root *TreeNode9, k int) {
	if root == nil {
		return
	}
	findK(root.left, k)
	a = a + 1
	if a == k {
		b = root.value
		return
	}
	findK(root.right, k)

}

func main20() {
	t10 := &TreeNode9{
		value: 10,
	}
	t8 := &TreeNode9{
		value: 8,
	}
	t9 := &TreeNode9{
		value: 9,
		left:  t8,
		right: t10,
	}

	t3 := &TreeNode9{
		value: 3,
	}
	t4 := &TreeNode9{
		value: 4,
		left:  t3,
	}
	t6 := &TreeNode9{
		value: 6,
	}
	t5 := &TreeNode9{
		value: 5,
		left:  t4,
		right: t6,
	}
	t7 := &TreeNode9{
		value: 7,
		left:  t5,
		right: t9,
	}

	t := t7

	findK(t, 6)
	fmt.Println(b)

}
