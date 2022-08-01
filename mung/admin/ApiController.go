package admin

type ApiController struct {
	BaseController
}

func (c *ApiController) NestPrepare(){

}

// Clear 清理服务器端缓存
func (c *ApiController) Clear(){
	c.JSONSuccess("服务端清理缓存成功")
}

// Workers 查询所有师傅的接口
func (c *ApiController) Workers(){

	username := c.GetString("username")
	phone := c.GetString("phone")
	page,_ := c.GetInt("page")
	lit,_ := c.GetInt("limit")

	// 查询所有user
	workers,err := c.Dao.QueryUserByArgs(username, phone, page, lit)
	if err != nil {
		c.Abort500(err)
	}
	// 根据条件得到user的总行数
	count, _ := c.Dao.QueryUserCount(username,phone)

	// 返回json
	c.JSONSuccess("", count, workers)

}


// Wpusers 查询所有师傅的接口
func (c *ApiController) Wpusers() {
	name := c.GetString("username")
	//phone := c.GetString("phone")
	page,_ := c.GetInt("page")
	lit,_ := c.GetInt("limit")

	// 查询所有user
	wpUsers,err := c.Dao.QueryWpUsersByArgs(name, page, lit)
	if err != nil {
		c.Abort500(err)
	}
	// 根据条件得到wpuser的总行数
	count, _ := c.Dao.QueryWpUserCountByArgs(name)

	// 返回json
	c.JSONSuccess("返回数据", count, wpUsers)

}