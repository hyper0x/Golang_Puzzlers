package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

// 方案1。
func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 方案2。
type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	// 方案1。
	x, y := 12, 23
	op := func(x, y int) int {
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)

	// 方案2。
	x, y = 56, 78
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
}
