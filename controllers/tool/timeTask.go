package tool

import (
	"github.com/astaxie/beego/toolbox"
	"beego-wechat/models"
	"fmt"
)

//func init()  {
//	shopTimeTask()
//}

//在第一个用户请求数据是就 触发一次
func ShopTimeTask()  {
	fmt.Println("shopTime in----")
	//timeStr := "0 */60 * * * *"
	timeStr := "0/30 * * * * *"

	getTokenTask := toolbox.NewTask("getToken",timeStr, func() error {
		err := models.GetAndUpdateAccessToken()
		if err != nil{
			//	请求access_token失败 业务处理
			fmt.Println("获取access_token失败!")
		}
		return nil
	})
	toolbox.AddTask("getToken",getTokenTask)
	toolbox.StartTask()
	fmt.Println("start----")
}