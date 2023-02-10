package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// 添字でアクセス
	fmt.Println(arr[1])

	// 要素数を省略
	names := [...]string{"yamada", "sato", "suzuki"}
	fmt.Println(names)

	// 配列の長さを取得
	fmt.Println(len(names))
}
