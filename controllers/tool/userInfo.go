package tool

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"bytes"
	"beego-wechat/models"
	"encoding/json"
	"beego-wechat/util"
)

type UserInfo struct {
	Subscribe	int	`json:"subscribe,omitempty"`
	Openid	string	`json:"openid,omitempty"`
	Nickname	string	`json:"nickname,omitempty"`
	Sex int	`json:"sex,omitempty"`
	Language string `json:"language,omitempty"`
	City	string	`json:"city,omitempty"`
	Province string	`json:"province,omitempty"`
	Country	string	`json:"country,omitempty"`
	Headimgurl	string	`json:"headimgurl,omitempty"`
	SubscribeTime int	`json:"subscribe_time,omitempty"`
	Unionid	string    `json:"unionid"`
	Remark	string	`json:"remark,omitempty"`
	Groupid	int	`json:"groupid,omitempty"`
	TagidList	[]int	`json:"tagid_list,omitempty"`
}

func GetWXUserInfo(openid string)(userInfo *models.WXUser, err error)  {
	token, err := models.GetAccessToken(1)
	if err != nil{
		fmt.Println("获取token报错")
		return
	}
	userInfoUrl := "https://api.weixin.qq.com/cgi-bin/user/info"
	requestLine := strings.Join([]string{userInfoUrl,
										 "?access_token=",token.AccessToken,
										 "&openid=",openid,
										 "&lang=zh_CN"},"")
	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode!= http.StatusOK{
		fmt.Println("发送url获取userinfo报错", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("发送get请求获取 wxUserInfo 读取返回body错误", err)
		return nil , err
	}
	if bytes.Contains(body, []byte("errcode")){
		ater := util.AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil{
			fmt.Printf("发送get请求获取 wxUserInfo 的错误信息 %+v\n", ater)
			return nil, err
		}
		return nil, fmt.Errorf("%s",ater.Errmsg)
	}else {
		err = json.Unmarshal(body, &userInfo)
		if err != nil{
			fmt.Println("发送get请求获取 wxUserInfo 返回数据json解析错误", err)
			return nil, err
		}
		return userInfo, nil
	}

}