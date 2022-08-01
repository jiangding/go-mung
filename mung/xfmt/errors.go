package xfmt

// 申明一个结构体
type errorString struct {
	s string
}

func New(txt string) error {
	return &errorString{txt }
}


// 实现了这个接口，就是自定义错误
func (e *errorString) Error() string {
	return e.s
}
