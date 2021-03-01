package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name int
	age  int
}

type PersonSlice []Person

func (v PersonSlice) Len() int {
	return len(v)
}

func (v PersonSlice) Less(i, j int) bool {
	if v[i].age == v[j].age {
		return v[i].name < v[j].name
	}
	return v[i].age > v[j].age
}

func (v PersonSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func main2() {
	s := []int{1, 2, 2, 3, 3, 3, 4, 4}

	// construct : num -> frequency
	myapp := make(map[int]int)
	for _, value := range s {
		if _, ok := myapp[value]; ok {
			myapp[value] = myapp[value] + 1
		} else {
			myapp[value] = 1
		}
	}

	// construct : {num, frequency}
	var persons []Person
	for key, value := range myapp {
		persons = append(persons, Person{
			name: key,
			age:  value,
		})
	}

	// sort slice({num, frequency})
	sort.Sort(PersonSlice(persons))
	for _, p := range persons {
		fmt.Println(p.name)
	}

	map1 := make(map[int]int)
	map1[1] = 1
	map1[2] = 2
	map1[3] = 3
	map1[4] = 3
	map1[5] = 5

	var keys []int
	for key, _ := range map1 {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

}
