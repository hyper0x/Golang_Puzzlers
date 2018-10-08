package main

import (
	"log"
	"sync"
	"time"
)

// counter 代表计数器。
type counter struct {
	num uint         // 计数。
	mu  sync.RWMutex // 读写锁。
}

// number 会返回当前的计数。
func (c *counter) number() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

// add 会增加计数器的值，并会返回增加后的计数。
func (c *counter) add(increment uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += increment
	return c.num
}

func main() {
	c := counter{}
	count(&c)
	redundantUnlock()
}

func count(c *counter) {
	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	go func() { // 用于增加计数。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= 20; j++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("The number in counter: %d [%d-%d]",
				c.number(), 1, j)
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= 20; k++ {
			time.Sleep(time.Millisecond * 300)
			log.Printf("The number in counter: %d [%d-%d]",
				c.number(), 2, k)
		}
	}()
	<-sign
	<-sign
	<-sign
}

func redundantUnlock() {
	var rwMu sync.RWMutex

	// 示例1。
	//rwMu.Unlock() // 这里会引发panic。

	// 示例2。
	//rwMu.RUnlock() // 这里会引发panic。

	// 示例3。
	rwMu.RLock()
	//rwMu.Unlock() // 这里会引发panic。
	rwMu.RUnlock()

	// 示例4。
	rwMu.Lock()
	//rwMu.RUnlock() // 这里会引发panic。
	rwMu.Unlock()
}
