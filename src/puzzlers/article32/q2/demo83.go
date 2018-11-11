package main

import (
	"io"
	"strings"
)

func main() {
	comment := "Because these interfaces and primitives wrap lower-level operations with various implementations, " +
		"unless otherwise informed clients should not assume they are safe for parallel execution."
	basicReader := strings.NewReader(comment)
	basicWriter := new(strings.Builder)

	// 示例1。
	reader1 := io.LimitReader(basicReader, 98)
	_ = interface{}(reader1).(io.Reader)

	// 示例2。
	reader2 := io.NewSectionReader(basicReader, 98, 89)
	_ = interface{}(reader2).(io.Reader)
	_ = interface{}(reader2).(io.ReaderAt)
	_ = interface{}(reader2).(io.Seeker)

	// 示例3。
	reader3 := io.TeeReader(basicReader, basicWriter)
	_ = interface{}(reader3).(io.Reader)

	// 示例4。
	reader4 := io.MultiReader(reader1)
	_ = interface{}(reader4).(io.Reader)

	// 示例5。
	writer1 := io.MultiWriter(basicWriter)
	_ = interface{}(writer1).(io.Writer)

	// 示例6。
	pReader, pWriter := io.Pipe()
	_ = interface{}(pReader).(io.Reader)
	_ = interface{}(pReader).(io.Closer)
	_ = interface{}(pWriter).(io.Writer)
	_ = interface{}(pWriter).(io.Closer)
}
