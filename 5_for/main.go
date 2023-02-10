package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 式のみの指定も可能
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// continue: 次のループへ, break: for分を抜ける
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println(i)
	}

	// rangeを利用した繰り返し処理
	names := []string{"yamada", "sato", "suzuki"}
	for index, name := range names {
		fmt.Println(index, name)
	}
}
