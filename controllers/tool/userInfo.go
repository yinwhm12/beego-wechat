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
	"github.com/astaxie/beego"
	"errors"
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

//get code
func GetCodeForAccessToken()  {
	//redirect_url := url.QueryEscape("http://yinwhm12.ngrok.cc/user_info")
	//urlStr := "https://open.weixin.qq.com/connect/oauth2/authorize?appid="+
	//	beego.AppConfig.String("appid")+"&redirect_uri="+ redirect_url+"response_type=code&scope=snsapi_base&"+
	//		"state=123#wechat_redirect"

}

type ResAccessToken struct {
	AccessToken string	`json:"access_token,omitempty"`
	ExpiresIn int64 `json:"expires_in,omitempty"`
	RefreshToken	string	`json:"refresh_token,omitempty"`
	Openid	string	`json:"openid,omitempty"`
	Scope	string	`json:"scope,omitempty"`
}

func GetOpenidByCode(code string)  error{
	requestLine := strings.Join([]string{"https://api.weixin.qq.com/sns/oauth2/access_token",
	"?appid=",beego.AppConfig.String("appid"),
	"&secret=",beego.AppConfig.String("secret"),
	"&code=",code,
	"&grant_type=authorization_code"},"")
	resp, err := http.Get(requestLine)
	if err != nil{
		fmt.Println("发送get请求获取openid错误",err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("statu err", resp.StatusCode)
		return errors.New("status")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		fmt.Println("读取openid时的body错", err)
		return err
	}
	if bytes.Contains(body, []byte("errcode")){
		ater := util.AccessTokenErrorResponse{}
		err = json.Unmarshal(body,ater)
		if err != nil{
			fmt.Printf("发送get请求获取 openid 的错误信息 %+v\n", ater)
			return  err
		}
	}else{
		rsepAccessToken := ResAccessToken{}
		err = json.Unmarshal(body, rsepAccessToken)
		if err != nil{
			fmt.Println("发送get请求获取 openid 返回数据json解析错误", err)
			return err
		}
	}
	return nil
}

//check access_token
func CheckAccessToken(accessToken, openID string)(b bool, err error)  {
	urlStr := fmt.Sprintf(beego.AppConfig.String("checkTokenURL"),accessToken,openID)
	var response []byte
	response, err = util.HttpGet(urlStr)
	if err != nil{
		return
	}
	var result util.AccessTokenErrorResponse
	err = json.Unmarshal(response, &result)
	if err != nil{
		return
	}
	if result.Errcode != 0{
		b = false
		return
	}
	b = true
	return
}

//get userInfo
func GetUserInformation(accessToken, openID string)(err error)  {
	urlStr :="https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken +
		"&openid=" + openID +
		"&lang=zh_CN"
	var response []byte
	response, err = util.HttpGet(urlStr)
	if err != nil{
		return
	}
	if bytes.Contains(response,[]byte("errcode")); err != nil {
		ater := util.AccessTokenErrorResponse{}
		err = json.Unmarshal(response, &ater)
		if err != nil{
			fmt.Printf("发送get请求获取 wxUserInfo 的错误信息 %+v\n", ater)
			return
		}
		//return nil, fmt.Errorf("%s",ater.Errmsg)
	}else{

	}
	return
}
