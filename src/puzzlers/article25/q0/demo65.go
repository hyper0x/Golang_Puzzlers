package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	coordinateWithChan()
	fmt.Println()
	coordinateWithWaitGroup()
}

func coordinateWithChan() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d [with chan struct{}]\n", num)
	max := int32(10)
	go addNum(&num, 1, max, func() {
		sign <- struct{}{}
	})
	go addNum(&num, 2, max, func() {
		sign <- struct{}{}
	})
	<-sign
	<-sign
}

func coordinateWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := int32(0)
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	max := int32(10)
	go addNum(&num, 3, max, wg.Done)
	go addNum(&num, 4, max, wg.Done)
	wg.Wait()
}

// addNum 用于原子地增加numP所指的变量的值。
func addNum(numP *int32, id, max int32, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		if currNum >= max {
			break
		}
		newNum := currNum + 2
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
		} else {
			fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}
