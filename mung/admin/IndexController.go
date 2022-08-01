package admin

type IndexController struct {
	BaseController
}

func (c *IndexController) NestPrepare(){
	c.MustLogin()
	// 设置分组
	c.Data["group"] = "grp1"
}

func (c *IndexController) Index(){

	c.Data["title"] = "首页"

	c.TplName = "admin/index.html"
}





