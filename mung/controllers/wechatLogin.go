package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"mung/admin"
	"mung/models"
	"time"
)

var Min *miniprogram.MiniProgram

func init() {
	// 微信
	wc := wechat.NewWechat()
	memory := cache.NewMemory()

	cfg := &miniConfig.Config{
		AppID:     "wx23ab2068f6132395",
		AppSecret: "89f3e176ed79e3cab1147c774d3d966d",
		Cache: memory,
	}
	Min = wc.GetMiniProgram(cfg)

	//utils.Display("a", a)
}

type WechatLogin struct {
	admin.BaseController
}

func (c *WechatLogin) Login() {
	// 根据前台返回的code
	code := c.GetString("code")
	//encryptedData := c.GetString("encryptedData")
	//iv := c.GetString("iv")
	userInfo := c.GetString("userInfo")

	// code 换取openid
	codeStr,_ := Min.GetAuth().Code2Session(code)

	// 解析json
	var wp models.Wpuser
	json.Unmarshal([]byte(userInfo), &wp)

	if err := c.Dao.SaveWpUserByStruct(admin.H{
		"sessionkey" : codeStr.SessionKey,
		"unionid" : codeStr.UnionID,
		"openid" : codeStr.OpenID,
		"nickname" : wp.Nickname,
		"avatar" : wp.Avatar,
		"gender" : wp.Gender,
		"language" : wp.Language,
		"city" : wp.City,
		"province" : wp.Province,
		"country" : wp.Country,
		"created_at":  time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		fmt.Printf(" wpUser添加出错了 %s\n", err)
	}

	c.Data["json"] = admin.H{"code":200, "msg":"授权成功!", "openid":codeStr.OpenID}
	c.ServeJSON()
}
