package main

import "fmt"

type IntColor struct {
	value int
	color bool // true red , false black
}

type UnLike struct {
	a IntColor
	b IntColor
}

type UnLikeInt struct {
	a int
	b int
}

func main10() {
	//values := []int{1, 2, 3, 4}

	ui1 := UnLikeInt{a: 1, b: 2}
	ui2 := UnLikeInt{a: 3, b: 4}
	ui3 := UnLikeInt{a: 1, b: 3}
	ui4 := UnLikeInt{a: 5, b: 6}
	ui5 := UnLikeInt{a: 6, b: 7}
	ui6 := UnLikeInt{a: 5, b: 7}

	valueMap := make(map[int]IntColor)
	var unLikeSlice []UnLike

	var unHandled []UnLikeInt

	for i, v := range []UnLikeInt{ui1, ui2, ui3, ui4, ui5, ui6} {
		if i == 0 {
			va := v.a
			vb := v.b
			ca := IntColor{value: va, color: true}
			cb := IntColor{value: vb, color: false}
			valueMap[va] = ca
			valueMap[vb] = cb
			unLikeSlice = append(unLikeSlice, UnLike{a: ca, b: cb})
		} else {
			va := v.a
			vb := v.b
			v0, ok := valueMap[va]
			v1, ok1 := valueMap[vb]
			if !ok && !ok1 {
				unHandled = append(unHandled, v)
				continue
			}
			var ca IntColor
			var cb IntColor
			if ok && !ok1 {
				ca = v0
				cb = IntColor{value: vb, color: !v0.color}
				unLikeSlice = append(unLikeSlice, UnLike{a: ca, b: cb})
			}
			if !ok && ok1 {
				cb = v1
				ca = IntColor{value: va, color: !cb.color}
				unLikeSlice = append(unLikeSlice, UnLike{a: ca, b: cb})
			}
			if ok && ok1 {
				cb = v1
				ca = v0
				unLikeSlice = append(unLikeSlice, UnLike{a: ca, b: cb})
			}
			valueMap[va] = ca
			valueMap[vb] = cb
		}
	}

	for _, u := range unLikeSlice {
		if u.a.color == u.b.color {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
}
