package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"time"
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

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, AppKey string, Nonce string, CurTime string, CheckSum string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	request, err2 := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err2 != nil {
		return ""
	}
	request.Header.Add("AppKey", AppKey)
	request.Header.Add("Nonce", Nonce)
	request.Header.Add("CurTime", CurTime)
	request.Header.Add("CheckSum", CheckSum)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}
