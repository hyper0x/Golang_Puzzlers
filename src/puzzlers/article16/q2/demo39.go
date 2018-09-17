package main

import (
	"fmt"
	//"time"
)

func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}
}
