package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"

	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"

	"mung/admin"
	_ "mung/routers"
	"os"
)
//const var init main
func main() {
	// 设置资源文件目录
	//beego.SetStaticPath("/","static/")

	// 注册自定义错误页面
	beego.ErrorController(&admin.ErrorController{})

	// 初始化日志
	initLog()

	// 开启task
	// utils.WarningDays()

	// 开启小程序
	// initProgram()

	beego.Run()

	// 关闭数据库
	//defer models.SqlDB.Close()

}

// 初始化日志输出目录
func initLog() {
	if err := os.MkdirAll("data/logs", 0777); err != nil {
		logs.Error(err)
		return
	}
	logs.SetLogger("file", `{"filename":"data/logs/mung.log","level":7,"maxlines":1000000,"daily":true,"maxdays":10}`)
	// 异步输出 设置缓冲chan大小
	logs.Async(1e3)
}

func initProgram(){

	wc := wechat.NewWechat()
	//redisOpts := &cache.RedisOpts{
	//	Host:        "127.0.0.1:6379",
	//	Database:    0,
	//	MaxActive:   10,
	//	MaxIdle:     10,
	//	IdleTimeout: 60, //second
	//}

	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Cache: memory,
	}
	miniprogram := wc.GetMiniProgram(cfg)
	miniprogram.GetAnalysis().GetAnalysisDailyRetain("2022-07-15", "2022-10-15")


}

