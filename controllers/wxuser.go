package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
)

type WXUserBaseControllers struct {
	beego.Controller
}

// @router / [get]
func (c *BaseControllers)GetUserInfo()  {
	fmt.Println("user_info 授权")
	fmt.Println(c.Ctx.Request.Body)
	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		fmt.Println("wrong body----")
		return
	}
	fmt.Println("body=",body)
}