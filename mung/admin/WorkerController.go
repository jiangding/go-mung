package admin

import (
	"mung/models"
)

type WorkerController struct {
	BaseController
}

func (c *WorkerController) NestPrepare(){
	c.MustLogin()
	// 设置分组
	c.Data["group"] = "grp1"
}

func (c *WorkerController) List(){

	c.Data["title"] = "师傅列表"

	c.TplName = "admin/worker_list.html"
}

// Add 添加展示页
func (c *WorkerController) Add(){

	c.Data["title"] = "添加"

	c.TplName = "admin/table/add.html"
}

// Edit 修改用户信息页
func (c *WorkerController) Edit(){

	uid,_ := c.GetInt("id")
	// 获取某一条user
	user, _ := c.Dao.QueryOneUserById(uid)

	c.Data["user"] = user
	c.TplName = "admin/table/worker_edit.html"
}


// ToDo 执行添加或者修改
func (c *WorkerController) ToDo(){
	uid,_ := c.GetInt("uid")
	username := c.GetMustString("username", "用户名不能为空!")
	phone := c.GetMustString("phone","手机不能为空!")
	status,_ := c.GetInt("status")
	account := c.GetMustString("account", "account不能为空")
	imgPath := c.GetString("img_path")

	if uid > 0 {
		// 执行修改
		if err := c.Dao.UpdateUserStructById(H{
			"nickname" : username,
			"account" : account,
			"phone" : phone,
			"status" : status,
			"avatar" : imgPath,
		}, uid); err != nil {
			//c.Abort500(errors.New("更新出错！"))
			c.JSONFailed("更新出错!")
		}
	}else {
		// 执行添加
		if err := c.Dao.SaveUser(&models.User{
			Nickname : username,
			Account : account,
			Phone : phone,
			Status : status,
			Avatar : imgPath,
		}); err != nil {
			//c.Abort500(errors.New("用户添加出错！"))
			c.JSONFailed("用户添加出错!")
		}
	}
	c.JSONSuccess("操作成功")
}

// Del 执行删除操作
func (c *WorkerController) Del(){
	uid,_ := c.GetInt("id")
	// 执行
	if err := c.Dao.DeleteUserById(uid); err != nil && uid > 0 {
		c.JSONFailed("删除出错")
	}else{
		c.JSONSuccess("删除成功")
	}
}











