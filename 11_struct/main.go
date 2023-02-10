package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func main() {
	// 初期化
	yamada := struct {
		Name string
		Age  int
	}{
		Name: "yamada",
		Age:  20,
	}
	fmt.Printf("%+v", yamada)

	// フィールドにアクセス
	yamada.Age = 30
	fmt.Println(yamada.Age)

	// typeを利用
	sato := Human{
		Name: "sato",
		Age:  20,
	}
	fmt.Printf("%+v", sato)
}
