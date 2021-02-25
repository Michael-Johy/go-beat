package main

import "fmt"
import "example.com/user/hello/morestrings"
import "github.com/google/go-cmp/cmp"

func main() {
	fmt.Print(morestrings.ReverseRunes("Hello1 World"))
	fmt.Print(cmp.Diff("Hello, World", "Hello, Go"))
}
