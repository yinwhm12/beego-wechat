package controllers

import (
	"github.com/astaxie/beego"
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"fmt"
	"beego-wechat/controllers/tool"
	"net/url"
)

type BaseControllers struct {
	beego.Controller
}

const Token  = "yinwhm12"
var once int = 1


// @router /sign [get]
func (c *BaseControllers)Sign()  {
	timestamp, nonce, signatureIn := c.GetString("timestamp"), c.GetString("nonce"), c.GetString("signature")
	signatureGen := makeSignature(timestamp,nonce)
	if signatureGen != signatureIn{
		fmt.Println("forbiddern")
		return
	}else{
		echostr := c.GetString("echostr")
		c.Ctx.WriteString(echostr)
		if once == 1{
			fmt.Println("first time ---token")
			tool.ShopTimeTask()
			tool.MenuTimeTask()
			once ++
		}
	}
	
}

func makeSignature(timestamp, nonce string)string  {
	s1 := []string{Token, timestamp, nonce}
	sort.Strings(s1)
	s := sha1.New()
	io.WriteString(s, strings.Join(s1,""))
	return fmt.Sprintf("%x",s.Sum(nil))
}


// @router / [get]
func (c *BaseControllers)Get()  {
	fmt.Println("user_info 授权")
	//fmt.Println(c.Ctx.Request.Body)
	//body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	//if err != nil{
	//	fmt.Println("wrong body----")
	//	return
	//}
	//fmt.Println("body=",body)
	if code := c.GetString("code"); code == ""{
		encodUrl := url.QueryEscape(beego.AppConfig.String("baseURL")+"/v1/wx_connect")
		urlStr := "https://open.weixin.qq.com/connect/oatuorize?appid="+
			beego.AppConfig.String("appid")+"&redirect_uri="+encodUrl+
			"&response_type=code&scope=snsapi_base&state=123#wechat_redirect"
		c.Redirect(urlStr,302)
	}else{

	}
}
