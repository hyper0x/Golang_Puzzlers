package op

import (
	"bytes"
	"encoding/json"
	"math/rand"
)

// box 代表数据盒子。
type box struct {
	Str   string
	Code  rune
	Bytes []byte
}

func MemProfile() error {
	max := 50000
	var buf bytes.Buffer
	for j := 0; j < max; j++ {
		seed := rand.Intn(95) + 32
		one := createBox(seed)
		b, err := genJSON(one)
		if err != nil {
			return err
		}
		buf.Write(b)
		buf.WriteByte('\t')
	}
	_ = buf.String()
	return nil
}

func createBox(seed int) box {
	if seed <= 0 {
		seed = 1
	}
	var array []byte
	size := seed * 8
	for i := 0; i < size; i++ {
		array = append(array, byte(seed))
	}
	return box{
		Str:   string(seed),
		Code:  rune(seed),
		Bytes: array,
	}
}

func genJSON(one box) ([]byte, error) {
	return json.Marshal(one)
}
