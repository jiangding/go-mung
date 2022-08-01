package admin

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
)

type PubController struct {
	BaseController
}

func (c *PubController) NestPrepare(){
	c.MustLogin()
	c.Data["group"] = "grp1"
}


// ChangePwd 修改密码
func (c *PubController) ChangePwd(){

	c.Data["title"] = "修改密码"

	c.TplName = "admin/change_pwd.html"
}

// ToChangePwd 执行修改
func (c *PubController) ToChangePwd(){
	// 获取信息
	old := strings.TrimSpace(c.GetString("old_password"))
	news := strings.TrimSpace(c.GetString("new_password"))
	again := strings.TrimSpace(c.GetString("again_password"))
	name := strings.TrimSpace(c.GetString("name"))

	if news != again {
		c.Abort500(errors.New("两次输入密码不一致"))
	}

	// 判断旧密码
	adm, err := c.Dao.QueryAdminByAccountAndPassword(name, old)
	if err != nil {
		c.JSONFailed("旧密码输入错误")
	}

	// 新密码
	// 获取md5加密后的密码
	h:= md5.New()
	h.Write([]byte(news))
	newpwd := hex.EncodeToString(h.Sum(nil))
	adm.Pwd = newpwd
	// 修改新密码
	if err :=  c.Dao.UpdatePwdByAccount(&adm, name); err != nil {
		c.JSONFailed("修改密码出错")
	}

	c.JSONSuccess("修改密码成功", 0, nil)
}


