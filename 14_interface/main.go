package main

import "fmt"

type Human interface {
	SayName()
}

type Man struct {
	Name string
}

func (m Man) SayName() {
	fmt.Println(m.Name)
}

type Woman struct {
	Name string
}

func main() {
	var man Human = Man{Name: "yamada"}
	man.SayName()

	// SayNameメソッドを実装していないためエラーになる
	// var woman Human = Woman{Name: "yamada"}
}
