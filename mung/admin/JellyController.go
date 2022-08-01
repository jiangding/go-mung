package admin

import (
	"errors"
	"fmt"
	"mung/jelly"
	"mung/xfmt"
)

type JellyController struct {
	BaseController
}

func (c *JellyController) NestPrepare() {
	// 判断是否登录咯
	//c.MustLogin()
}

func (c *JellyController) Index() {

	// 版本1
	//j := jelly.NewJelly(10,"tiy",18, "12049393848",nil)

	// 可选参数版本, 有就有，没有就会使用默认的
	//j:= jelly.NewJelly(111)
	// 扩展性好
	j := jelly.NewJelly(10, jelly.WithName("jackson"), jelly.WithPhone("110"))

	fmt.Println("123")

	err1 := errors.New("hello1")
	err2 := xfmt.New("hello2")
	err3 := fmt.Errorf("hell3 %v", err2)

	fmt.Printf("%T %v \n", err1, err1)
	fmt.Printf("%T %v \n", err2, err2)
	fmt.Printf("%T %v \n", err3, err3)

	c.JSONSuccess("jelly", 1, j)

}
