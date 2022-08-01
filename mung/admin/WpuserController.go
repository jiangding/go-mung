package admin

type WpuserController struct {
	BaseController
}

// 结构体前面* 表示指向结构体的指针, 传递的是引用地址
// &符号的意思是对变量取地址
// *符号的意思是对指针取值
// &是取指针，*是取值


func (c *WpuserController) NestPrepare(){
	// 必须登录
	c.MustLogin()
	// 设置分组
	c.Data["group"] = "grp1"
}

// List 首页展示
func (c *WpuserController) List(){

	c.Data["title"] = "小程序用户"

	c.TplName = "admin/wpuser_list.html"
}

// Edit 修改用户信息页
func (c *WpuserController) Edit(){

	uid,_ := c.GetInt("id")
	// 获取某一条user
	wpuser, _ := c.Dao.QueryOneWpUserById(uid)

	c.Data["user"] = wpuser.Nickname
	c.TplName = "admin/table/wpuser_edit.html"
}