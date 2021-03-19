package main

import "fmt"

type NTree struct {
	value    int
	intB     bool
	children []NTree
}

var result []int

func travel(root NTree) {
	if root.intB {
		result = append(result, root.value)
		return
	}
	for _, children := range root.children {
		travel(children)
	}
}

func main21() {
	t1 := NTree{
		value: 1,
		intB:  true,
	}
	c1 := NTree{
		value:    -1,
		intB:     false,
		children: []NTree{t1, t1},
	}

	c2 := NTree{
		value: 2,
		intB:  true,
	}

	c3 := NTree{
		value:    -1,
		intB:     false,
		children: []NTree{t1, t1},
	}

	root := NTree{
		value:    -1,
		intB:     false,
		children: []NTree{c1, c2, c3},
	}

	travel(root)

	fmt.Println(result)
}
