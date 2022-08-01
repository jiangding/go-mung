package admin

import (
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/server/web/captcha"
	"mung/models"
	"strings"
)
// 验证码
var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	// 设置验证码参数
	cpt.ChallengeNums = 3
	cpt.StdWidth = 120
	cpt.StdHeight = 50

}


type LoginController struct {
	BaseController
}

func (c *LoginController) Login(){

	c.TplName = "admin/login.html"
}
// ToLogin 执行登录
func (c *LoginController) ToLogin(){
	// 获取信息
	name := strings.TrimSpace(c.GetString("name"))
	pwd := strings.TrimSpace(c.GetString("pwd"))
	//captcha := c.GetString("captcha")

	// 验证码
	b := cpt.VerifyReq(c.Ctx.Request)

	if !b {
		c.JSONFailed("验证码出错")
		return
	}

	var (
		adm models.Admin
		err error
	)
	if adm, err = c.Dao.QueryAdminByAccountAndPassword(name, pwd); err != nil {
		// 登录失败
		c.JSONFailed("账号或密码错误")

	}else{

		//将保存到session
		c.SetSession(SESSION_ADMIN_KEY, adm)
		// 保存到cookie中
		c.Ctx.SetCookie(SESSION_ADMIN_KEY,adm.Account,3600*24*2, ".abc.com")

		c.JSONSuccess("Loading...")
	}

}

// Logout 执行退出
func (c *LoginController) Logout() {

	c.DelSession(SESSION_ADMIN_KEY)
	// 清除cookie 让值为空就能达到删除效果
	c.Ctx.SetCookie(SESSION_ADMIN_KEY,"",3600*24*2,".abc.com")

	c.JSONSuccess("退出成功")
}