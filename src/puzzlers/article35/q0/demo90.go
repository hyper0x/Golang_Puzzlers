package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd1, err := syscall.Socket(
		syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Printf("socket error: %v\n", err)
		return
	}
	defer syscall.Close(fd1)
	fmt.Printf("The file descriptor of socket：%d\n", fd1)

	// 省略若干代码。
	// 如果真要完全使用syscall包中的程序实体建立网络连接的话，
	// 过程太过繁琐而且完全没有必要。
	// 所以，我在这里就不做展示了。
}
