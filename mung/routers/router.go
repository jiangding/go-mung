package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mung/admin"
	"mung/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	// 匹配数据，没有为空
	// beego.Router("/admin/index/?:id", &admin.IndexController{})
	// 匹配数字，定义类型
	beego.Router("/admin/index/:id:string", &admin.IndexController{})
	//
	// beego.Router("/api/:id([0-9]+)", &admin.IndexController{})

	// 首页
	beego.Router("/admin", &admin.IndexController{}, "get:Index")

	// 师傅列表
	beego.Router("/admin/list", &admin.WorkerController{}, "get:List")
	// 添加页面add
	beego.Router("/admin/worker_add", &admin.WorkerController{}, "get:Add")
	// 修改页面edit
	beego.Router("/admin/worker_edit", &admin.WorkerController{}, "get:Edit")
	// 执行worker添加和修改
	beego.Router("/admin/worker2do", &admin.WorkerController{}, "post:ToDo")
	// 删除worker
	beego.Router("/admin/worker_del", &admin.WorkerController{}, "post:Del")


	// 小程序用户列表
	beego.Router("/admin/wpuser_list", &admin.WpuserController{}, "get:List")
	beego.Router("/admin/wpuser_edit", &admin.WpuserController{}, "get:Edit")



	// hello
	beego.Router("/admin/hello", &admin.HelloController{},"get:Hello")
	beego.Router("/admin/hello2", &admin.HelloController{},"get:Hello2")


	// 后台首页Api
	beego.Router("/adminapi/clear",&admin.ApiController{},"get:Clear")
	// worker列表api
	beego.Router("/adminapi/workers",&admin.ApiController{},"get:Workers")
	// wpuser列表api
	beego.Router("/adminapi/wpusers",&admin.ApiController{},"get:Wpusers")


}
