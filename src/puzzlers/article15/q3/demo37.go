package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))

	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Printf("nameP == &(dogP.name)? %v\n",
		nameP == &(dogP.name))
	fmt.Printf("The name of dog is %q.\n", *nameP)

	*nameP = "monster"
	fmt.Printf("The name of dog is %q.\n", dogP.name)
	fmt.Println()

	// 示例2。
	// 下面这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期。
	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)

}
