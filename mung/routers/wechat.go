package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mung/controllers"
)

func init() {
	// 微信小程序登录
	beego.Router("/api/login",&controllers.WechatLogin{},"get:Login")

	//
	//beego.Router("", &controllers.WechatLogin{}, "get:")
	//
	//beego.Router("", &controllers.WechatLogin{}, "get:")
}
