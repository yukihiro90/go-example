package main

import "fmt"

func main() {
	name := "yamada"
	switch name {
	case "sato":
		fmt.Println("名前はsatoです")
	case "yamada", "suzuki":
		fmt.Println("名前はyamadaまたはsuzukiです")
	default:
		fmt.Println("名前はこれら以外です")
	}

	// case1, 2まで処理される
	i := 1
	switch i {
	case 1:
		fmt.Println("値は1です")
		fallthrough
	case 2:
		fmt.Println("値は2です")
	case 3:
		fmt.Println("値は1です")
	}
}
