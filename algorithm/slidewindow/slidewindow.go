package main

//
//func main() {
//
//}
//
//func find(s string, t string) {
//	sa := []rune(s)
//	ta := []rune(t)
//
//	tm := make(map[rune]int)
//	for _, item := range ta {
//		if count, ok := tm[item]; ok {
//			tm[item] = count + 1
//		} else {
//			tm[item] = 1
//		}
//	}
//
//	resultL := 0
//	resultR := 0
//
//	var result []string
//
//	wm := make(map[rune]int)
//	valid := 0
//
//	left := 0
//	right := 0
//
//	for {
//		if right < len(sa) {
//			rv := sa[right]
//			right++
//
//			if ct, ok := tm[rv]; ok {
//				if cw, ok1 := wm[rv]; ok1 {
//					wm[rv] = cw + 1
//				} else {
//					wm[rv] = 1
//				}
//				if wm[rv] == ct {
//					valid++
//				}
//			}
//
//			for {
//				if valid == len(tm) {
//					resultL = left
//					resultR = right
//					lv := sa[left]
//					left++
//					if _, ok := tm[lv]; ok {
//						wm[lv]--
//						if wm[lv] != tm[lv] {
//							valid--
//						}
//					}
//				}
//			}
//
//		}
//	}
//
//}
