package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")

	defer func() {
		fmt.Println("Enter defer function.")

		// recover函数的正确用法。
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}

		fmt.Println("Exit defer function.")
	}()

	// recover函数的错误用法。
	fmt.Printf("no panic: %v\n", recover())

	// 引发panic。
	panic(errors.New("something wrong"))

	// recover函数的错误用法。
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")
}
