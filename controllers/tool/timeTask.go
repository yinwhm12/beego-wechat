package tool

import (
	"github.com/astaxie/beego/toolbox"
	"beego-wechat/models"
)

//在第一个用户请求数据是就 触发一次
func shopTimeTask()  {
	timeStr := "0 */60 * * * *"
	getTokenTask := toolbox.NewTask("getToken",timeStr, func() error {
		err := models.GetAndUpdateAccessToken()
		if err != nil{
			//	请求access_token失败 业务处理
		}
		return nil
	})
	toolbox.AddTask("getToken",getTokenTask)
	toolbox.StartTask()
}