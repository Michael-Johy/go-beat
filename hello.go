package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"
)

import "example.com/user/hello/morestrings"
import "example.com/user/hello/model"
import "example.com/user/hello/file"
import "github.com/google/go-cmp/cmp"

func main() {
	fmt.Print(morestrings.ReverseRunes("Hello1 World"))
	fmt.Print(cmp.Diff("Hello, World", "Hello, Go"))

	reverseCase := morestrings.ReverseCase

	ss := "a n c"
	fmt.Printf("Before = %q\n", ss)
	fmt.Println(strings.Map(reverseCase, ss))

	// test struct
	location := model.Location{X: 11.11, Y: 22.22}

	student1 := &model.Student{Name: "yang", Lct: location}

	stuJson, err := json.Marshal(student1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stuJson))

	student2 := model.Student{}
	_ = json.Unmarshal(stuJson, &student2)
	fmt.Println(student2.Name)

	//control structures
	arr := []string{"1", "2"}
	b := len(arr)
	if a := b; a != 1 {
		fmt.Println("a != 1")
	}

	aaa, bbb := aa()
	fmt.Println(aaa)
	fmt.Println(bbb)

	// file
	content, err := file.Read("D:\\aa")
	if nil != err {
		log.Fatal("读取错误")
	}
	fmt.Println(content)

	mymap := map[string]int{
		"yan1": 1,
		"yan2": 2,
		"yan3": 3,
	}
	age, ok := mymap["yan1"]
	fmt.Println(age)
	fmt.Println(ok)

	fmt.Println(math.Exp2(2))

}

func aa() (a, b int32) {
	a = 1
	b = 2
	return
}

type Stringer interface {
	String() (n int, err error)
}

type Joiner interface {
	Join() (n int, err error)
}

func testType(a interface{}) {
	switch value := a.(type) {
	case int:
		fmt.Print("int", value)
	case float32:
		fmt.Print("float32", value)
	default:
		fmt.Printf("have type %T", value)
	}
}
