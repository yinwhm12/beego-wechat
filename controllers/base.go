package controllers

import (
	"github.com/astaxie/beego"
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"fmt"
	"beego-wechat/controllers/tool"
)

type BaseControllers struct {
	beego.Controller
}

const Token  = "yinwhm12"
var once int = 1

// @router / [get]
func (c *BaseControllers)Get()  {
	timestamp, nonce, signatureIn := c.GetString("timestamp"), c.GetString("nonce"), c.GetString("signature")
	signatureGen := makeSignature(timestamp,nonce)
	if signatureGen != signatureIn{
		fmt.Println("forbiddern")
		return
	}else{
		echostr := c.GetString("echostr")
		c.Ctx.WriteString(echostr)
		if once == 1{
			tool.ShopTimeTask()
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
