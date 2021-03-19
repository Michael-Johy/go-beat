package main

import "fmt"

// 位运算 均已计算机内部数据的二进制格式进行计算
// 正数：原码 = 反码 = 补码
// 负数：原码 -> 符号位不变，数值位取反 => 反码
//      反码 -> 反码 + 1 => 补码
func main9() {
	fmt.Println(-1 | 2)
	fmt.Println(-1 & 2)
	fmt.Println(3 ^ -2)

	negetive()
	positive()
	negPos()

}

//判断2个数都是负数
func negetive() {
	fmt.Println("neg ")
	fmt.Println((-1 & -1) < 0)

	fmt.Println((-1 & 1) < 0)
	fmt.Println((1 & 1) < 0)
}

//判断2个都是正数
func positive() {
	fmt.Println("pos ")
	fmt.Println((1 | 1) > 0)

	fmt.Println((1 | -1) > 0)
	fmt.Println((-1 | -1) > 0)
}

//一正一负
func negPos() {
	fmt.Println("^ ")
	fmt.Println((1 ^ -1) < 0)

	fmt.Println((1 ^ 1) < 0)
	fmt.Println((-1 ^ -1) < 0)
}
