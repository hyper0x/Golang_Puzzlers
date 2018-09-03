package main

import "fmt"

func main() {
	// 示例1。
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
	fmt.Println()

	// 示例2。
	slice1 := []string{"x", "y", "z"}
	fmt.Printf("The slice: %v\n", slice1)
	slice2 := modifySlice(slice1)
	fmt.Printf("The modified slice: %v\n", slice2)
	fmt.Printf("The original slice: %v\n", slice1)
	fmt.Println()

	// 示例3。
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The modified complex array: %v\n", complexArray2)
	fmt.Printf("The original complex array: %v\n", complexArray1)
}

// 示例1。
func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

// 示例2。
func modifySlice(a []string) []string {
	a[1] = "i"
	return a
}

// 示例3。
func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"
	a[2] = []string{"o", "p", "q"}
	return a
}
