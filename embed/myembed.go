package main

import (
	_ "embed"
	"fmt"
)

//go:embed txt/hello
var s string

func main() {
	fmt.Println(s)
}
