package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	// 第二个衍生问题的示例。
	num := uint32(18)
	fmt.Printf("The number: %d\n", num)
	delta := int32(-3)
	atomic.AddUint32(&num, uint32(delta))
	fmt.Printf("The number: %d\n", num)
	atomic.AddUint32(&num, ^uint32(-(-3)-1))
	fmt.Printf("The number: %d\n", num)

	fmt.Printf("The two's complement of %d: %b\n",
		delta, uint32(delta)) // -3的补码。
	fmt.Printf("The equivalent: %b\n", ^uint32(-(-3)-1)) // 与-3的补码相同。
	fmt.Println()

	// 第三个衍生问题的示例。
	forAndCAS1()
	fmt.Println()
	forAndCAS2()
}

// forAndCAS1 用于展示简易的自旋锁。
func forAndCAS1() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	go func() { // 定时增加num的值。
		defer func() {
			sign <- struct{}{}
		}()
		for {
			time.Sleep(time.Millisecond * 500)
			newNum := atomic.AddInt32(&num, 2)
			fmt.Printf("The number: %d\n", newNum)
			if newNum == 10 {
				break
			}
		}
	}()
	go func() { // 定时检查num的值，如果等于10就将其归零。
		defer func() {
			sign <- struct{}{}
		}()
		for {
			if atomic.CompareAndSwapInt32(&num, 10, 0) {
				fmt.Println("The number has gone to zero.")
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	<-sign
	<-sign
}

// forAndCAS2 用于展示一种简易的（且更加宽松的）互斥锁的模拟。
func forAndCAS2() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	max := int32(20)
	go func(id int, max int32) { // 定时增加num的值。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; ; i++ {
			currNum := atomic.LoadInt32(&num)
			if currNum >= max {
				break
			}
			newNum := currNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, currNum, newNum) {
				fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
			}
		}
	}(1, max)
	go func(id int, max int32) { // 定时增加num的值。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 0; ; j++ {
			currNum := atomic.LoadInt32(&num)
			if currNum >= max {
				break
			}
			newNum := currNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, currNum, newNum) {
				fmt.Printf("The number: %d [%d-%d]\n", newNum, id, j)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, j)
			}
		}
	}(2, max)
	<-sign
	<-sign
}
