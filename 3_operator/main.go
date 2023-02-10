package main

import "fmt"

func main() {
	// - 演算子

	// 加算
	res1 := 1 + 2
	fmt.Println(res1)

	// 減算
	res2 := 2 - 1
	fmt.Println(res2)

	// 掛け算
	res3 := 3 * 2
	fmt.Println(res3)

	// 割り算
	res4 := 4 / 2
	fmt.Println(res4)

	res5 := 5 / 2
	fmt.Println(res5) // 2になる

	res6 := 5.0 / 2
	fmt.Println(res6) // 2になる

	// 余り
	res7 := 8 / 3
	fmt.Println(res7)

	// - 代入演算子
	i := 1

	i += 2
	fmt.Println(i)

	i -= 1
	fmt.Println(i)

	i++
	fmt.Println(i)

	i--
	fmt.Println(i)
}
