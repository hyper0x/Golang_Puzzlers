package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

func main() {
	// buffer 代表缓冲区。
	var buffer bytes.Buffer

	const (
		max1 = 5  // 代表启用的goroutine的数量。
		max2 = 10 // 代表每个goroutine需要写入的数据块的数量。
		max3 = 10 // 代表每个数据块中需要有多少个重复的数字。
	)

	// mu 代表以下流程要使用的互斥锁。
	var mu sync.Mutex
	// protecting 用于表明是否真正使用互斥锁。
	// 改变该变量的值，运行程序，并观察程序打印的内容。
	protecting := true
	// sign 代表信号的通道。
	sign := make(chan struct{}, 2)

	for i := 1; i <= max1; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= max2; j++ {
				// 准备数据。
				header := fmt.Sprintf("\n[id: %d, iteration: %d]",
					id, j)
				data := fmt.Sprintf(" %d", id*j)
				// 写入数据。
				if protecting {
					mu.Lock()
				}
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error: %s [%d]", err, id)
				}
				for k := 0; k < max3; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error: %s [%d]", err, id)
					}
				}
				if protecting {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}

	<-sign
	<-sign
	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	log.Printf("The contents:\n%s", data)
}
