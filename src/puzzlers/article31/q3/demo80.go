package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 示例1。
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap())
	fmt.Println()

	unreadBytes := buffer1.Bytes()
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Println()

	contents = "cdefg"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()

	// 只要扩充一下之前拿到的未读字节切片unreadBytes，
	// 就可以用它来读取甚至修改buffer中的后续内容。
	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Println()

	value := byte('X')
	fmt.Printf("Set a byte in the unread bytes to %v ...\n", value)
	unreadBytes[len(unreadBytes)-2] = value
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer1.Bytes())
	fmt.Println()

	// 不过，在buffer的内容容器真正扩容之后就无法这么做了。
	contents = "hijklmn"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()

	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Print("\n\n")

	// 示例2。
	// Next方法返回的后续字节切片也存在相同的问题。
	contents = "12"
	buffer2 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer2.Cap())
	fmt.Println()

	nextBytes := buffer2.Next(2)
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
	fmt.Println()

	contents = "34567"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer2.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer2.Cap())
	fmt.Println()

	// 只要扩充一下之前拿到的后续字节切片nextBytes，
	// 就可以用它来读取甚至修改buffer中的后续内容。
	nextBytes = nextBytes[:cap(nextBytes)]
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
	fmt.Println()

	value = byte('X')
	fmt.Printf("Set a byte in the next bytes to %v ...\n", value)
	nextBytes[len(nextBytes)-2] = value
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer2.Bytes())
	fmt.Println()

	// 不过，在buffer的内容容器真正扩容之后就无法这么做了。
	contents = "89101112"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer2.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer2.Cap())
	fmt.Println()

	nextBytes = nextBytes[:cap(nextBytes)]
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
}
