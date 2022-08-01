package admin

type StuffController struct {
	BaseController
}

func (c *StuffController) NestPrepare(){
	// 判断是否登录咯
	c.MustLogin()
}

func (c *StuffController) Index(){
	c.Data["group"] = "grp1"


	//


	c.TplName = "admin/study/stuff.html"
}
