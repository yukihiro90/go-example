package main

import "fmt"

func main() {
	name := "yamada"
	if name == "yamada" {
		fmt.Println("名前はyamadaです。")
	}

	i := 2
	if i%2 == 0 {
		fmt.Println("値は偶数です。")
	} else {
		fmt.Println("値は奇数です。")
	}

	if i != 0 && i != 1 {
		fmt.Println("値は0でも1でもありません。")
	}

	if i == 0 || i == 2 {
		fmt.Println("値は0または2です。")
	}

	if age := 20; age >= 20 {
		fmt.Println("年齢は20歳以上です。")
	}
}
