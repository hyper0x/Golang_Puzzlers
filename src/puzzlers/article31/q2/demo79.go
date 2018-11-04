package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 示例1。
	var contents string
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The length of new buffer with contents %q: %d\n",
		contents, buffer1.Len())
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap())
	fmt.Println()

	contents = "12345"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()

	contents = "67"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()

	contents = "89"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Print("\n\n")

	// 示例2。
	contents = "abcdefghijk"
	buffer2 := bytes.NewBufferString(contents)
	fmt.Printf("The length of new buffer with contents %q: %d\n",
		contents, buffer2.Len())
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer2.Cap())
	fmt.Println()

	n := 10
	fmt.Printf("Grow the buffer with %d ...\n", n)
	buffer2.Grow(n)
	fmt.Printf("The length of buffer: %d\n", buffer2.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer2.Cap())
	fmt.Print("\n\n")

	// 示例3。
	var buffer3 bytes.Buffer
	fmt.Printf("The length of new buffer: %d\n", buffer3.Len())
	fmt.Printf("The capacity of new buffer: %d\n", buffer3.Cap())
	fmt.Println()

	contents = "xyz"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer3.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer3.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer3.Cap())
}
