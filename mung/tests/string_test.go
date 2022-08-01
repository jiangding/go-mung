package test

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// go test -bench="String$" -benchmem .

// 1. 直接拼接
// 在go语言中，字符串（string类型）是不可变的（如果我们需要改变字符串，就要转成byte类型，再转回string类型），
// 因此字符串之间的拼接实际上是创建了一个新的字符串，就会 不断开辟内存空间 ，如果频繁进行字符串拼接，就会对性能产生严重的影响。
func BenchmarkString(b *testing.B) {
	// 创建字符串切片
	elems := make([]string, 100000, 100000)
	for i := 0; i < len(elems); i++ {
		elems[i] = strconv.Itoa(i)
	}
	sum := ""
	length := len(elems)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			sum += elems[j]
		}
	}
	b.StopTimer()
}

// 2. fmt.Sprintf("%s",xxxxx) Go语言 fmt.Sprintf 的底层实现中，用到了反射的机制。那为什么反射会慢呢？
func BenchmarkFmtSprintfString(b *testing.B) {
	elems := make([]string, 100000, 100000)
	for i := 0; i < len(elems); i++ {
		elems[i] = strconv.Itoa(i)
	}
	length := len(elems)
	sum := ""
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			sum += fmt.Sprintf("%s", elems[j])
		}
	}
	b.StopTimer()
}

// 3. string.Builder  strings.Builder  bytes.Buffer，包括切片 []byte 的内存 是以倍数申请的。
func BenchmarkBuilderString(b *testing.B) {
	elems := make([]string, 100000, 100000)
	for i := 0; i < len(elems); i++ {
		elems[i] = strconv.Itoa(i)
	}
	var builder strings.Builder
	length := len(elems)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			builder.WriteString(elems[j])
		}
	}
	b.StopTimer()
}

// 4. bytes.Builder
func BenchmarkByteBufferString(b *testing.B) {
	elems := make([]string, 100000, 100000)
	for i := 0; i < len(elems); i++ {
		elems[i] = strconv.Itoa(i)
	}
	buffer := new(bytes.Buffer)
	length := len(elems)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			buffer.WriteString(elems[j])
		}
	}
	b.StopTimer()
}

// 5. byte 拼接  而第五个，毫无疑问事先分配好了内存，肯定是最快的了。
func BenchmarkByteConcatString(b *testing.B) {
	elems := make([]string, 100000, 100000)
	for i := 0; i < len(elems); i++ {
		elems[i] = strconv.Itoa(i)
	}
	length := len(elems)
	buf := make([]byte, 0, len(elems))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			buf = append(buf, elems[j]...)
		}
	}
	b.StopTimer()
}



