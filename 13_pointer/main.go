package main

import "fmt"

type Human struct {
	Age int
}

func (h *Human) addAge() {
	h.Age += 1
}

func set(i *int) {
	*i = 5
}

func main() {
	// 型を定義
	var i int
	var ip *int // ゼロ値はnill
	fmt.Println(i, ip)

	// アドレスを代入
	// ip = i  // 型が違うのでエラーになる
	ip = &i // 変数iのアドレスをipに代入
	fmt.Println(ip)

	// 値の代入や取得には*をつける
	*ip = 1
	fmt.Println(*ip)

	set(&i)
	fmt.Println(i) // 5になる

	// ポインタの利用例
	human := Human{Age: 10}
	human.addAge()
	fmt.Println(human.Age)
}
