package main

import "fmt"

// 変数定義①: 型を指定して代入
var country string = "japan"

// 変数定義②: 型を指定せずに代入
var state = "tokyo"

func main() {
	// 変数定義③: var無しで宣言（関数内のみで使用可能）
	city := "sinjuku"
	fmt.Println(country, state, city)

	// まとめて宣言する方法
	var (
		name = "yamada"
		age  = 30
	)
	fmt.Println(name, age)

	// ゼロ値
	var count int
	var text string
	var flg bool
	fmt.Println(count, text, flg)
}
