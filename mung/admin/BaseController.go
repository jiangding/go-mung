package admin

import (
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"mung/models"
)

// SESSION_ADMIN_KEY session中的key值
const SESSION_ADMIN_KEY = "SESSION_ADMIN_KEY"

type H map[string]interface{} // 公用的map

type BaseController struct {
	beego.Controller
	IsLogin bool //标识 用户是否登陆
	Admin    models.Admin //登陆的用户
	Dao   *models.DB
}

// NestPreparer 如果子controller 实现了NestPrepare()方法，就实现了该接口
type NestPreparer interface {
	NestPrepare()
}

// Prepare 预处理方法
func (c *BaseController) Prepare() {
	// 获取当前path
	c.Data["Path"] = c.Ctx.Request.RequestURI
	// 初始化DB
	c.Dao = models.NewDB()

	// 验证用户是否登陆，判断session中是否存在用户，存在就已经登陆，不存在就没有登陆。
	c.IsLogin = false
	tu := c.GetSession(SESSION_ADMIN_KEY)

	if tu != nil { // session有的话就取出来
		// 断言: 是否实现了期望的接口或者具体的类型
		if u, ok := tu.(models.Admin); ok {
			c.Admin = u
			c.Data["admin"] = u.Account
			c.IsLogin = true
		}
	}
	//
	tt := c.Ctx.GetCookie(SESSION_ADMIN_KEY)

	if tt != "" {
		c.Data["admin"] = tt
		c.IsLogin = true
	}

	// 断言: 判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法。
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}

	// 没有拿到就跳转去登录把
	//if !c.IsLogin {
	//	c.Ctx.Redirect(302, "/admin/login")
	//}

	c.Data["IsLogin"] = c.IsLogin
}

// GetMustString 必须的获取数据
func (c *BaseController) GetMustString(key string, msg string) string {
	keys := c.GetString(key, "")
	if len(keys) == 0 {
		c.Abort500(errors.New(msg))
	}
	return keys
}

// MustLogin 一定要登录
func (c *BaseController) MustLogin(){
	// 判断是否登录咯
	if !c.IsLogin {
		c.Ctx.Redirect(302, "/admin/login")
	}
}



// Abort500 封装500错误
func (c *BaseController) Abort500(err error) {
	c.Data["error"] = err
	c.Abort("500")
}

type ReturnJson struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	//Action string      `json:"action,omitempty"`
	Count  int64         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
func (c *BaseController) JSONSuccess(msg string, args ...interface{}) {
	var count int64
	var data interface{}
	// 遍历可选参数
	for _, arg := range args {
		switch arg.(type) { // 断言
		case int64:
			count = arg.(int64)
		case interface{}:
			data = arg.(interface{})
		}
	}

	c.Data["json"] = &ReturnJson{
		Code:   0,
		Msg:    msg,
		Count: count,
		Data: data,
	}
	c.ServeJSON()
}

func (c *BaseController) JSONFailed(msg string) {

	c.Data["json"] = &ReturnJson{
		Code:   500,
		Msg:    msg,

	}
	c.ServeJSON()
}