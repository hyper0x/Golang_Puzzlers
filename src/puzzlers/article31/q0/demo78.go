package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 示例1。
	var buffer1 bytes.Buffer
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()

	// 示例2。
	p1 := make([]byte, 7)
	n, _ := buffer1.Read(p1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
}
