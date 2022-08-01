package admin

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"path"
	"time"
)

type UploadController struct {
	beego.Controller
}


// UploadPic 图片上传
func (c *UploadController) UploadPic(){
	//获取图片信息
	f, h, err := c.GetFile("file")
	if err != nil {
		log.Fatal("获取文件错误 ", err)
	}
	//清除临时文件
	defer f.Close()
	//获取文件后缀
	var fileSuffix string
	fileSuffix = path.Ext(h.Filename)
	//获取当前时间戳
	timeNow := time.Now().Unix()

	//设置保存后的路径
	imageUrl := "static/upload/"  + fmt.Sprintf("%d", timeNow) + fileSuffix
	//保存文件
	c.SaveToFile("file", imageUrl) // 保存位置在 static/upload, 没有文件夹要先创建

	c.Data["json"] = H{"code": 200, "img": "/"+imageUrl ,"message": "上传成功"}
	c.ServeJSON()
	return
}
