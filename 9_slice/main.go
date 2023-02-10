package main

import "fmt"

func main() {
	// 基本の定義方法
	arr := [...]int{1, 2, 3, 4, 5}
	arr1 := arr[2:4]
	fmt.Println(arr1)

	// スライスの変更が配列に反映される
	arr1[0] = 6
	fmt.Println(arr)

	// 配列の変更がスライスに反映される
	arr[3] = 7
	fmt.Println(arr1)

	// 可変長配列として利用
	i := []int{0, 1, 2}

	// 要素を追加
	i = append(i, 3, 4)
	fmt.Println(i)
}
