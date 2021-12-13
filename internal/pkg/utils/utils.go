package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strconv"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func MD5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}

func RandomNumStr(digit int) string {
	var numList = make([]int64, 0)
	for i := 0; i < digit; i++ {
		result, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			numList = append(numList, 0)
		} else {
			numList = append(numList, result.Int64())
		}
	}
	var bf bytes.Buffer
	for _, num := range numList {
		randomNumStr := strconv.FormatInt(num, 10)
		bf.WriteString(randomNumStr)
	}
	return bf.String()
}
