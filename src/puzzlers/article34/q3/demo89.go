package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type argDesc struct {
	action string
	flag   int
	perm   os.FileMode
}

func main() {
	// 示例1。
	fmt.Printf("The mode for dir:\n%32b\n", os.ModeDir)
	fmt.Printf("The mode for named pipe:\n%32b\n", os.ModeNamedPipe)
	fmt.Printf("The mode for all of the irregular files:\n%32b\n", os.ModeType)
	fmt.Printf("The mode for permissions:\n%32b\n", os.ModePerm)
	fmt.Println()

	// 示例2。
	fileName1 := "something3.txt"
	filePath1 := filepath.Join(os.TempDir(), fileName1)
	fmt.Printf("The file path: %s\n", filePath1)

	argDescList := []argDesc{
		{
			"Create",
			os.O_RDWR | os.O_CREATE,
			0644,
		},
		{
			"Reuse",
			os.O_RDWR | os.O_TRUNC,
			0666,
		},
		{
			"Open",
			os.O_RDWR | os.O_APPEND,
			0777,
		},
	}

	defer os.Remove(filePath1)
	for _, v := range argDescList {
		fmt.Printf("%s the file with perm %o ...\n", v.action, v.perm)
		file1, err := os.OpenFile(filePath1, v.flag, v.perm)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		info1, err := file1.Stat()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("The file permissions: %o\n", info1.Mode().Perm())
	}
}
