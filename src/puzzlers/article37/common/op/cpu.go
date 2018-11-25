package op

import (
	"bytes"
	"math/rand"
	"strconv"
)

func CPUProfile() error {
	max := 10000000
	var buf bytes.Buffer
	for j := 0; j < max; j++ {
		num := rand.Int63n(int64(max))
		str := strconv.FormatInt(num, 10)
		buf.WriteString(str)
	}
	_ = buf.String()
	return nil
}
