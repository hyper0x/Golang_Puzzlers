package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
)

// ioTypes 代表了io代码包中的所有接口的反射类型。
var ioTypes = []reflect.Type{
	reflect.TypeOf((*io.Reader)(nil)).Elem(),
	reflect.TypeOf((*io.Writer)(nil)).Elem(),
	reflect.TypeOf((*io.Closer)(nil)).Elem(),

	reflect.TypeOf((*io.ByteReader)(nil)).Elem(),
	reflect.TypeOf((*io.RuneReader)(nil)).Elem(),
	reflect.TypeOf((*io.ReaderAt)(nil)).Elem(),
	reflect.TypeOf((*io.Seeker)(nil)).Elem(),
	reflect.TypeOf((*io.WriterTo)(nil)).Elem(),
	reflect.TypeOf((*io.ByteWriter)(nil)).Elem(),
	reflect.TypeOf((*io.WriterAt)(nil)).Elem(),
	reflect.TypeOf((*io.ReaderFrom)(nil)).Elem(),

	reflect.TypeOf((*io.ByteScanner)(nil)).Elem(),
	reflect.TypeOf((*io.RuneScanner)(nil)).Elem(),
	reflect.TypeOf((*io.ReadSeeker)(nil)).Elem(),
	reflect.TypeOf((*io.ReadCloser)(nil)).Elem(),
	reflect.TypeOf((*io.WriteCloser)(nil)).Elem(),
	reflect.TypeOf((*io.WriteSeeker)(nil)).Elem(),
	reflect.TypeOf((*io.ReadWriter)(nil)).Elem(),
	reflect.TypeOf((*io.ReadWriteSeeker)(nil)).Elem(),
	reflect.TypeOf((*io.ReadWriteCloser)(nil)).Elem(),
}

func main() {
	// 示例1。
	file1 := (*os.File)(nil)
	fileType := reflect.TypeOf(file1)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Type %T implements\n", file1)
	for _, t := range ioTypes {
		if fileType.Implements(t) {
			buf.WriteString(t.String())
			buf.WriteByte(',')
			buf.WriteByte('\n')
		}
	}
	output := buf.Bytes()
	output[len(output)-2] = '.'
	fmt.Printf("%s\n", output)

	// 示例2。
	fileName1 := "something1.txt"
	filePath1 := filepath.Join(os.TempDir(), fileName1)
	var paths []string
	paths = append(paths, filePath1)
	dir, _ := os.Getwd()
	paths = append(paths, filepath.Join(dir[:len(dir)-1], fileName1))
	for _, path := range paths {
		fmt.Printf("Create a file with path %s ...\n", path)
		_, err := os.Create(path)
		if err != nil {
			var underlyingErr string
			if _, ok := err.(*os.PathError); ok {
				underlyingErr = "(path error)"
			}
			fmt.Printf("error: %v %s\n", err, underlyingErr)
			continue
		}
		fmt.Println("The file has been created.")
	}
	fmt.Println()

	// 示例3。
	fmt.Println("New a file associated with stderr ...")
	file3 := os.NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	if file3 != nil {
		file3.WriteString(
			"The Go language program writes something to stderr.\n")
	}
	fmt.Println()

	// 示例4。
	fmt.Printf("Open a file with path %s ...\n", filePath1)
	file4, err := os.Open(filePath1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println("Write something to the file ...")
	_, err = file4.WriteString("something")
	var underlyingErr string
	if _, ok := err.(*os.PathError); ok {
		underlyingErr = "(path error)"
	}
	fmt.Printf("error: %v %s\n", err, underlyingErr)
	fmt.Println()

	// 示例5。
	fmt.Printf("Open a file with path %s ...\n", filePath1)
	file5a, err := os.Open(filePath1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf(
		"Is there only one file descriptor for the same file in the same process? %v\n",
		file5a.Fd() == file4.Fd())
	file5b := os.NewFile(file5a.Fd(), filePath1)
	fmt.Printf("Can the same file descriptor represent the same file? %v\n",
		file5b.Name() == file5a.Name())
	fmt.Println()

	// 示例6。
	fmt.Printf("Reuse a file on path %s ...\n", filePath1)
	file6, err := os.OpenFile(filePath1, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	contents := "something"
	fmt.Printf("Write %q to the file ...\n", contents)
	n, err := file6.WriteString(contents)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("The number of bytes written is %d.\n", n)
	}
}
