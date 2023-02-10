package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
}
