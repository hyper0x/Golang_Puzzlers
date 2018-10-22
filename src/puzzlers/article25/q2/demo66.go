package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 示例1。
	var counter uint32
	var once sync.Once
	once.Do(func() {
		atomic.AddUint32(&counter, 1)
	})
	fmt.Printf("The counter: %d\n", counter)
	once.Do(func() {
		atomic.AddUint32(&counter, 2)
	})
	fmt.Printf("The counter: %d\n", counter)
	fmt.Println()

	// 示例2。
	once = sync.Once{}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		once.Do(func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("Do task. [1-%d]\n", i)
				time.Sleep(time.Second)
			}
		})
		fmt.Println("Done. [1]")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task. [2]")
		})
		fmt.Println("Done. [2]")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task. [3]")
		})
		fmt.Println("Done. [3]")
	}()
	wg.Wait()
	fmt.Println()

	// 示例3。
	once = sync.Once{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("fatal error: %v\n", p)
			}
		}()
		once.Do(func() {
			fmt.Println("Do task. [4]")
			panic(errors.New("something wrong"))
			fmt.Println("Done. [4]")
		})
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task. [5]")
		})
		fmt.Println("Done. [5]")
	}()
	wg.Wait()
}
