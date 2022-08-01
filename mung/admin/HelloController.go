package admin

type HelloController struct {
	BaseController
}

func (c *HelloController) NestPrepare(){
	c.MustLogin()
}


func (c *HelloController) Hello(){
	c.Data["group"] = "grp1"
	c.Data["title"] = "hello world"

	c.TplName = "admin/hello.html"
}

func (c *HelloController) Hello2(){
	c.Data["group"] = "grp2"
	c.Data["title"] = "hello world222"

	c.TplName = "admin/hello2.html"
}
