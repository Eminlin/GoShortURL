package common

import (
	"math"
	"strings"
)

var chars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//Uint32ToB62 转62进制
func Uint32ToB62(num uint32) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

//Uint64ToB62 转62进制
func Uint64ToB62(num uint64) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

//B62Decode 解析
func B62Decode(str string) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(chars, str[i])
		num += int64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
