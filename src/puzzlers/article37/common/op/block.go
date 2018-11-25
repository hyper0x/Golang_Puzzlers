package op

import (
	"math/rand"
	"sync"
	"time"
)

func BlockProfile() error {
	max := 100
	senderNum := max / 2
	receiverNum := max / 4
	ch1 := make(chan int, max/4)

	var senderGroup sync.WaitGroup
	senderGroup.Add(senderNum)
	repeat := 50000
	for j := 0; j < senderNum; j++ {
		go send(ch1, &senderGroup, repeat)
	}

	go func() {
		senderGroup.Wait()
		close(ch1)
	}()

	var receiverGroup sync.WaitGroup
	receiverGroup.Add(receiverNum)
	for j := 0; j < receiverNum; j++ {
		go receive(ch1, &receiverGroup)
	}
	receiverGroup.Wait()
	return nil
}

func send(ch1 chan int, wg *sync.WaitGroup, repeat int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 10)
	for k := 0; k < repeat; k++ {
		elem := rand.Intn(repeat)
		ch1 <- elem
	}
}

func receive(ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for elem := range ch1 {
		_ = elem
	}
}
