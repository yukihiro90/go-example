package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

// ① メソッド定義
func (h Human) SayName() {
	fmt.Println(h.Name)
}

// ② レシーバーは受け取らなくても良い
func (Human) Hello() {
	fmt.Println("hello")
}

// ③ 他のメソッドを呼び出すこともできる
func (h Human) Say() {
	h.Hello()
	h.SayName()
}

// ④ 元データの変更
func (h *Human) SetName(name string) {
	h.Name = name
}

func main() {
	yamada := Human{
		Name: "yamada",
		Age:  20,
	}

	yamada.SayName()

	yamada.Hello()

	yamada.Say()

	yamada.SetName("sato")
	yamada.SayName()
}
