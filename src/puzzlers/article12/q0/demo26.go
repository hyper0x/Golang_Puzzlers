package main

import "fmt"

type Printer func(content string) (n int, err error)

func printToStd(content string) (bytesNum int, err error) {
	return fmt.Println(content)
}

func main() {
	var p Printer
	p = printToStd
	p("something")
}
