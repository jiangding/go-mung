package admin

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	c.TplName="admin/error/404.html"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "500服务器内部出错了"
	c.TplName="admin/error/500.html"
}
