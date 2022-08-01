package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

// RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//Substring 获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || start > end {
		return ""
	}

	if end > length {
		end = length
	}

	if start == 0 && end >= length {
		return source
	}

	return string(r[start:end])
}

func SubstrAfter(s, substr string) string {
	a := strings.Index(s, substr)
	return Substring(s, a+len(substr), len(s))
}

func SubstrBefore(s, substr string) string {
	b := strings.Index(s, substr)
	return Substring(s, 0, b)
}

func SubstrBetween(s, afterStr, beforeStr string) string {
	a := strings.Index(s, afterStr)
	b := strings.Index(s, beforeStr)
	return Substring(s, a+len(afterStr), b)
}

func ArrayContains(arr []string, s string) bool {
	for _, t := range arr {
		if t == s {
			return true
		}
	}
	return false
}

// MbStrLen 中文字符串长度
func MbStrLen(str string) int {
	return utf8.RuneCountInString(str)
}

// Str2bytes 字符串转byte
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2str byte转字符串
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// JsonToMap json字符串bytes了下转map数据类型
func JsonToMap(str string) map[string]interface{} {

	jsonMap := map[string]interface{}{}

	json.Unmarshal(Str2bytes(str), &jsonMap)
	//fmt.Println("Str2bytes:", Str2bytes(jstr), "-end")
	fmt.Println("获取到的json=", jsonMap)
	return jsonMap
}
