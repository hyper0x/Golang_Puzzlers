package main

import (
	"fmt"
	"net"
	"time"
)

type dailArgs struct {
	network string
	address string
	timeout time.Duration
}

func main() {
	dialArgsList := []dailArgs{
		{
			"tcp",
			"google.cn:80",
			time.Millisecond * 500,
		},
		{
			"tcp",
			"google.com:80",
			time.Second * 2,
		},
		{
			// 如果在这种情况下发生的错误是：
			// "connect: operation timed out"，
			// 那么代表着什么呢？
			//
			// 简单来说，此错误表示底层的socket在连接网络服务的时候先超时了。
			// 这时抛出的其实是'syscall.ETIMEDOUT'常量代表的错误值。
			"tcp",
			"google.com:80",
			time.Minute * 4,
		},
	}
	for _, args := range dialArgsList {
		fmt.Printf("Dial %q with network %q and timeout %s ...\n",
			args.address, args.network, args.timeout)
		ts1 := time.Now()
		conn, err := net.DialTimeout(args.network, args.address, args.timeout)
		ts2 := time.Now()
		fmt.Printf("Elapsed time: %s\n", time.Duration(ts2.Sub(ts1)))
		if err != nil {
			fmt.Printf("dial error: %v\n", err)
			fmt.Println()
			continue
		}
		defer conn.Close()
		fmt.Printf("The local address: %s\n", conn.LocalAddr())
		fmt.Printf("The remote address: %s\n", conn.RemoteAddr())
		fmt.Println()
	}
}
