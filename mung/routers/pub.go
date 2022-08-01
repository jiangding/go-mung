package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mung/admin"
)

func init() {

	// 上传图片api
	beego.Router("/admin/uploadPic",&admin.UploadController{},"post:UploadPic")

	// stuff
	beego.Router("/admin/stuff",&admin.StuffController{},"get:Index")

	// jelly
	beego.Router("/admin/jelly",&admin.JellyController{},"get:Index")


	// 登录页
	beego.Router("/admin/login", &admin.LoginController{}, "get:Login")
	beego.Router("/admin/to_login", &admin.LoginController{}, "post:ToLogin")
	beego.Router("/admin/logout", &admin.LoginController{}, "post:Logout")


	// 修改密码
	beego.Router("/admin/change_pwd", &admin.PubController{}, "get:ChangePwd")
	beego.Router("/admin/to_change_pwd", &admin.PubController{}, "post:ToChangePwd")

}
