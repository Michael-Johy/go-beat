package main

import "fmt"

func main23() {
	s := [][]int{{1, 2}, {2, 3}, {3, 4}}

	var s1 [6]int

	for i, v := range s {
		for j, v1 := range v {
			s1[i*2+j] = v1
		}
	}
	fmt.Println(s1)

}
