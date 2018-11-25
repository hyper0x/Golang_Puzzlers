package main

import (
	"fmt"
	"net/http"
)

func main() {
	host := "google.cn"

	// 示例1。
	url1 := "http://" + host
	fmt.Printf("Send request to %q with method GET ...\n", url1)
	resp1, err := http.Get(url1)
	if err != nil {
		fmt.Printf("request sending error: %v\n", err)
		return
	}
	defer resp1.Body.Close()
	line1 := resp1.Proto + " " + resp1.Status
	fmt.Printf("The first line of response:\n%s\n", line1)
	fmt.Println()

	// 示例2。
	url2 := "https://golang." + host
	fmt.Printf("Send request to %q with method GET ...\n", url2)
	var httpClient1 http.Client
	resp2, err := httpClient1.Get(url2)
	if err != nil {
		fmt.Printf("request sending error: %v\n", err)
		return
	}
	defer resp2.Body.Close()
	line2 := resp2.Proto + " " + resp2.Status
	fmt.Printf("The first line of response:\n%s\n", line2)
}
