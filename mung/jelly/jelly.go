package jelly

// Jelly 结构体
type Jelly struct {
	Id    int
	Name  string
	Age   int
	Phone string

	Db interface{}
}

//func NewJelly(id int, name string, age int, phone string, db interface{}) *Jelly {
//	// 这个写法乍看之下是没啥问题的，
//	// 但是一旦Foo结构内部的字段进行了变化，增加或者减少了，
//	// 这个初始化函数NewJelly就怎么看怎么别扭了。参数继续增加？那么所有调用方的地方也都需要进行修改
//	return &Jelly{
//		Id: id,
//		Name: name,
//		Age: age,
//		Phone: phone,
//		Db: db,
//	}
//}

// JeyOption 使用Option写法来进行改造
type JeyOption func(jelly *Jelly) //函数类型, 将jelly指针传递进去，能让内部函数进行修改

// WithName 代表Name为可选参数
func WithName(name string) JeyOption {
	return func(jelly *Jelly) {
		jelly.Name = name
	}
}

// WithAge 代表Age为可选参数
func WithAge(age int) JeyOption {
	return func(jelly *Jelly) {
		jelly.Age = age
	}
}

// WithPhone 代表Phone为可选参数
func WithPhone(phone string) JeyOption {
	return func(jelly *Jelly) {
		jelly.Phone = phone
	}
}

// WithDB 代表db为可选参数
func WithDB(db interface{}) JeyOption {
	return func(jelly *Jelly) {
		jelly.Db = db
	}
}

// NewJelly 初始化方法
func NewJelly(id int, options ...JeyOption) *Jelly {
	// 默认值
	jelly := &Jelly{
		Id:    id,
		Name:  "defalut",
		Age:   10,
		Phone: "13666868888",
		Db:    nil,
	}
	// 如果有设置就循环设置，扩展性好
	for _, option := range options {
		option(jelly)
	}

	return jelly
}
