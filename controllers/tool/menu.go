package tool

import (
	"fmt"
	"net/http"
	"strings"
	"bytes"
	"beego-wechat/models"
)

func PushMenuCreate(accessToken string, menuJsonBytes []byte)error  {
	postReq, err := http.NewRequest("POST",
		strings.Join([]string{" https://api.weixin.qq.com/cgi-bin/menu/create?access_token=",accessToken},""),
		bytes.NewReader(menuJsonBytes))

	if err != nil{
		fmt.Println("向微信发送建立建立菜单失败",err)
		return err
	}
	postReq.Header.Set("Constent-Type","application/json;encoding=utf-8")
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil{
		fmt.Println("client向微信发送菜单建立失败",err)
		return err
	}else {
		fmt.Println("client 发送成功")
	}
	defer resp.Body.Close()
	return nil
}

func createMenu()  {
	menuStr := `{
            "button": [
            {
                "name": "进入商城",
                "type": "view",
                "url": "http://www.baidu.com/"
            },
            {

                "name":"管理中心",
                 "sub_button":[
                        {
                        "name": "用户中心",
                        "type": "click",
                        "key": "molan_user_center"
                        },
                        {
                        "name": "公告",
                        "type": "click",
                        "key": "molan_institution"
                        }]
            },
            {
                "name": "资料修改",
                "type": "view",
                "url": "http://www.baidu.com/user_view"
            }
            ]
        }`
	token, err := models.GetAccessToken(1)
	if err != nil{
		fmt.Println("从数据库获取token失败",err)
		return
	}
	PushMenuCreate(token.AccessToken, []byte(menuStr))

}