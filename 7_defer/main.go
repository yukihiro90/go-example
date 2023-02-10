package main

import "fmt"

func main() {
	defer fmt.Println("4番目の処理")
	defer fmt.Println("3番目の処理")
	defer fmt.Println("2番目の処理")

	fmt.Println("1番目の処理")
}
