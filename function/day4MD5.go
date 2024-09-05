package functions

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	// "golang.org/x/text/number"
)

func MD5(input string) int {
	number := 0
	for {
		code := input + strconv.Itoa(number)
		// md5.Sum that takes a []byte and return a [16]byte array
		hash := md5.Sum(StrByte(code))
		hashex := hex.EncodeToString((hash[:]))

		if hashex[:5] == "00000" {
			return number
		}
		number++
	}
}

func StrByte(s string) []byte {
	b := []byte{}
	for i := 0; i < len(s); i++ {
		b = append(b, s[i])
	}
	return b
}

func MD5b(input string) int {
	number := 0
	for {
		code := input + strconv.Itoa(number)
		// md5.Sum that takes a []byte and return a [16]byte array
		hash := md5.Sum(StrByte(code))
		hashex := hex.EncodeToString((hash[:]))

		if hashex[:6] == "000000" {
			return number
		}
		number++
	}
}
