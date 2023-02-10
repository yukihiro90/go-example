package main

import "fmt"

func main() {
	value := map[string]int{"a": 1, "b": 2}
	fmt.Println(value)

	// キーを指定して取得
	fmt.Println(value["a"])

	// 2つ目に存在するかのbool値が返ってくる
	val, exist := value["b"]
	fmt.Println(val, exist) // 2, true

	// 存在確認
	if _, exist := value["c"]; exist {
		fmt.Println("cの要素は存在します。")
	} else {
		fmt.Println("cの要素は存在しません。")
	}

	// 削除
	delete(value, "b")
	fmt.Println(value)
}
