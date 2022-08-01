package xfmt

import (
	"io"
	"os"
	"reflect"
	"sync"
	"unicode/utf8"
)

const (
	commaSpaceString  = ", "
	nilAngleString    = "<nil>"
	nilParenString    = "(nil)"
	nilString         = "nil"

)

func Println(a ...any) (n int, err error) {
	return XFprintln(os.Stdout, a...)
}

func XFprintln(w io.Writer, a ...any) (n int, err error) {
	p := newPrinter()
	p.doPrintln(a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func newPrinter() *pp {
	p := ppFree.Get().(*pp)  //类型断言 就是将接口类型的值(x)，转换成类型(T)。格式为：x.(T)
	p.panicking = false
	p.erroring = false
	p.wrapErrs = false
	return p
}

func (p *pp) doPrintln(a ...any) {
	for argNum, arg := range a {
		if argNum > 0 {
			p.buf.writeByte(' ')
		}
		p.printArg(arg, 'v')
	}
	p.buf.writeByte('\n')
}

func (p *pp) printArg(arg any, verb rune){
	p.arg = arg
	p.value = reflect.Value{}

	if arg == nil {
		switch verb {
		case 'T', 'v':
			p.fmt.padString(nilAngleString)
		default:
			//p.badVerb(verb)
		}
		return
	}
}

// Use simple []byte instead of bytes.Buffer to avoid large dependency.
type buffer []byte

func (b *buffer) write(p []byte) {
	*b = append(*b, p...)
}

func (b *buffer) writeString(s string) {
	*b = append(*b, s...)
}

func (b *buffer) writeByte(c byte) {
	*b = append(*b, c)
}

func (bp *buffer) writeRune(r rune) {
	if r < utf8.RuneSelf {
		*bp = append(*bp, byte(r))
		return
	}

	b := *bp
	n := len(b)
	for n+utf8.UTFMax > cap(b) {
		b = append(b, 0)
	}
	w := utf8.EncodeRune(b[n:n+utf8.UTFMax], r)
	*bp = b[:n+w]
}


// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.
type pp struct {
	buf buffer

	// arg holds the current item, as an interface{}.
	arg any

	// value is used instead of arg for reflect values.
	value reflect.Value

	// fmt is used to format basic items such as integers or strings.
	fmt fmt

	// reordered records whether the format string used argument reordering.
	reordered bool
	// goodArgNum records whether the most recent reordering directive was valid.
	goodArgNum bool
	// panicking is set by catchPanic to avoid infinite panic, recover, panic, ... recursion.
	panicking bool
	// erroring is set when printing an error string to guard against calling handleMethods.
	erroring bool
	// wrapErrs is set when the format string may contain a %w verb.
	wrapErrs bool
	// wrappedErr records the target of the %w verb.
	wrappedErr error
}
// 一个异步池
// 那么这个池子的目的是什么呢？其实一句话就可以说明白，就是为了复用已经使用过的对象，来达到优化内存使用和回收的目的
// 如果不够了呢，自己会通过new产生一些，当你放回去了之后这些对象会被别人进行复用
// 当对象特别大并且使用非常频繁的时候可以大大的减少对象的创建和回收的时间
var ppFree = sync.Pool{
	New: func() any { return new(pp) },
}

// free saves used pp structs in ppFree; avoids an allocation per invocation.
func (p *pp) free() {
	// Proper usage of a sync.Pool requires each entry to have approximately
	// the same memory cost. To obtain this property when the stored type
	// contains a variably-sized buffer, we add a hard limit on the maximum buffer
	// to place back in the pool.
	//
	// See https://golang.org/issue/23199
	if cap(p.buf) > 64<<10 {
		return
	}

	p.buf = p.buf[:0]
	p.arg = nil
	p.value = reflect.Value{}
	p.wrappedErr = nil
	ppFree.Put(p)
}


