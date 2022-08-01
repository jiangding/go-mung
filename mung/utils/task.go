package utils

import (
	"github.com/beego/beego/v2/adapter/toolbox"

)

func sycTask() error {

	// logs.Error("定时任务MyTask2 hello time:%v", time.Now())

	return nil
}


func WarningDays(){
	// 1.
	tk1 := toolbox.NewTask("tk1", "0/10 * * * * *", sycTask)

	// 2. 可以测试开启运行
	//err := tk1.Run()
	//if err != nil {
	//	//t.Fatal(err)
	//}
	// 3.
	toolbox.AddTask("tk1", tk1)

	// 4.
	toolbox.StartTask()
	//defer toolbox.StopTask()
}




