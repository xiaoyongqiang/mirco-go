package utils

import (
	"math/rand"
	"time"
	"fmt"
)

var (
	codes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen = len(codes)
)

func RandStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}

	return string(data)
}

func RandNumStr(size int) string {
	if size < 0 {
		return "-1"
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randStr := fmt.Sprintf("%v", rnd.Int31n(int32(pow(10, size))))
	for i := 0; size - len(randStr) > i; i++ {
		randStr = "0" + randStr
	}
	return randStr
}

func pow(x, n int) int {
	ret := 1
	for n != 0 {
		if n % 2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}