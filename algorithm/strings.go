package main

import "fmt"

func main12() {
	i := 1
	i++
	fmt.Println(i)
	fmt.Println(check("level"))
}

func check(s string) bool {
	r := []rune(s)

	for i, j := 0, len(r)-1; i <= j; {
		if r[i] == r[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
