package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func echo(x, y int) (int, int) {
	return x, y
}

func main() {
	// 基本の使い方
	res := sum(1, 2)
	fmt.Println(res)

	// 複数の戻り値
	x, y := echo(1, 2)
	fmt.Println(x, y)

	// 無名関数
	name := "yamada"
	func() {
		fmt.Println(name)
	}()

	// 関数を変数に代入
	print := func() {
		fmt.Println(name)
	}
	print()
}
