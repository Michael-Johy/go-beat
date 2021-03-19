package main

import (
	"errors"
	"fmt"
	"math"
)

type IntSlice []int

// remove E  shiftDown or shiftUp
func (v *IntSlice) remove(value int) error {
	target := -1
	for i, item := range *v {
		if value == item {
			target = i
		}
	}
	if target == -1 {
		return errors.New("not exist")
	}

	(*v)[target], (*v)[len(*v)-1] = (*v)[len(*v)-1], (*v)[target]

	*v = (*v)[:len(*v)-1]

	//try shiftUp
	//parent index
	pi := int(math.Ceil(float64(target)/2 - 1))
	for {
		if pi > -1 && (*v)[pi] < (*v)[target] {
			(*v)[pi], (*v)[target] = (*v)[target], (*v)[pi]
			target = pi
			pi = int(math.Ceil(float64(target)/2 - 1))
		} else {
			break
		}
	}

	// try shiftDown
	// children index
	ci := target*2 + 1
	ci2 := target*2 + 2
	for {
		if ci < len(*v) && (*v)[ci] > (*v)[target] {
			(*v)[ci], (*v)[target] = (*v)[target], (*v)[ci]
			target = ci
			ci = target*2 + 1
			ci2 = target*2 + 1
			continue
		} else if ci2 < len(*v) && (*v)[ci2] > (*v)[target] {
			(*v)[ci2], (*v)[target] = (*v)[target], (*v)[ci2]
			target = ci2
			ci = target*2 + 1
			ci2 = target*2 + 1
			continue
		} else {
			break
		}
	}
	fmt.Println(*v)
	return nil
}

// insert shiftUp
func (v *IntSlice) insert(value int) error {

	*v = append(*v, value)

	target := len(*v) - 1

	//try shiftUp
	//parent index
	pi := int(math.Ceil(float64(target)/2 - 1))
	for {
		if pi > -1 && (*v)[pi] < (*v)[target] {
			(*v)[pi], (*v)[target] = (*v)[target], (*v)[pi]
			target = pi
			pi = int(math.Ceil(float64(target)/2 - 1))
		} else {
			break
		}
	}
	fmt.Println(*v)
	return nil
}

// remove E  shiftDown or shiftUp
func (v *IntSlice) removeIndex(target int) error {
	if target < 0 || target > len(*v)-1 {
		return errors.New("index not valid")
	}

	(*v)[target], (*v)[len(*v)-1] = (*v)[len(*v)-1], (*v)[target]

	*v = (*v)[:len(*v)-1]

	//try shiftUp
	//parent index
	pi := int(math.Ceil(float64(target)/2 - 1))
	for {
		if pi > -1 && (*v)[pi] < (*v)[target] {
			(*v)[pi], (*v)[target] = (*v)[target], (*v)[pi]
			target = pi
			pi = int(math.Ceil(float64(target)/2 - 1))
		} else {
			break
		}
	}

	// try shiftDown
	// children index
	ci := target*2 + 1
	ci2 := target*2 + 2
	for {
		if ci < len(*v) && (*v)[ci] > (*v)[target] {
			(*v)[ci], (*v)[target] = (*v)[target], (*v)[ci]
			target = ci
			ci = target*2 + 1
			ci2 = target*2 + 1
			continue
		} else if ci2 < len(*v) && (*v)[ci2] > (*v)[target] {
			(*v)[ci2], (*v)[target] = (*v)[target], (*v)[ci2]
			target = ci2
			ci = target*2 + 1
			ci2 = target*2 + 1
			continue
		} else {
			break
		}
	}
	fmt.Println(*v)
	return nil

}

func main7() {
	heap := IntSlice{10, 7, 9, 5, 4, 3, 6}
	//heap.remove(10)
	heap.removeIndex(0)

	heap1 := IntSlice{10, 7, 9, 5, 4, 3, 6}
	heap1.insert(11)

}
