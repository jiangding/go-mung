package utils

import (
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

func scy(){

	req := httplib.Get("http://beego.me/")
	_, err := req.String()
	if err != nil {
		logs.Error(err)
	}

}
