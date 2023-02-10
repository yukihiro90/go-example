package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// - ① 並列処理のサンプル
	// 複数の並列処理の順序は保証されない
	go func() {
		fmt.Println("① 出力2-1")
	}()
	go func() {
		fmt.Println("① 出力2-2")
	}()

	fmt.Println("① 出力1")
	time.Sleep(1 * time.Second)
	fmt.Println("① 出力3")

	// - ② 並列処理の完了を待って次の処理を行う場合
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("② 出力1-1")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("② 出力1-2")
	}()
	wg.Wait()
	fmt.Println("② 出力3")
}
