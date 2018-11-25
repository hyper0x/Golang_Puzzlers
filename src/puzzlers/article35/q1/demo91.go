package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"runtime"
)

func main() {
	network := "tcp"
	host := "google.cn"
	reqStrTpl := `HEAD / HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Host: %s
User-Agent: Dialer/%s



`

	// 示例1。
	network1 := network + "4"
	address1 := host + ":80"
	fmt.Printf("Dial %q with network %q ...\n", address1, network1)
	conn1, err := net.Dial(network1, address1)
	if err != nil {
		fmt.Printf("dial error: %v\n", err)
		return
	}
	defer conn1.Close()

	reqStr1 := fmt.Sprintf(reqStrTpl, host, runtime.Version())
	fmt.Printf("The request:\n%s\n", reqStr1)
	_, err = io.WriteString(conn1, reqStr1)
	if err != nil {
		fmt.Printf("write error: %v\n", err)
		return
	}
	fmt.Println()

	reader1 := bufio.NewReader(conn1)
	line1, err := reader1.ReadString('\n')
	if err != nil {
		fmt.Printf("read error: %v\n", err)
		return
	}
	fmt.Printf("The first line of response:\n%s\n", line1)
	fmt.Println()

	// 示例2。
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS10,
	}
	network2 := network
	address2 := host + ":443"
	fmt.Printf("Dial %q with network %q ...\n", address2, network2)
	conn2, err := tls.Dial(network2, address2, tlsConf)
	if err != nil {
		fmt.Printf("dial error: %v\n", err)
		return
	}
	defer conn2.Close()

	reqStr2 := fmt.Sprintf(reqStrTpl, host, runtime.Version())
	fmt.Printf("The request:\n%s\n", reqStr2)
	_, err = io.WriteString(conn2, reqStr2)
	if err != nil {
		fmt.Printf("write error: %v\n", err)
		return
	}

	reader2 := bufio.NewReader(conn2)
	line2, err := reader2.ReadString('\n')
	if err != nil {
		fmt.Printf("read error: %v\n", err)
		return
	}
	fmt.Printf("The first line of response:\n%s\n", line2)
	fmt.Println()
}
