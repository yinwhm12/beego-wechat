package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"beego-wechat/util"
	"github.com/astaxie/beego"
)

type AccessToken struct {
	Id int `orm:"pk;column(id);auto"`
	AccessToken string	`orm:"column(id)"`
}

func (at *AccessToken)TableName() string {
	return "token"
}

func init()  {
	orm.RegisterModel(new(AccessToken))
}

func GetAndUpdateAccessToken()error  {
	o := orm.NewOrm()
	token := AccessToken{Id:1}
	o.ReadOrCreate(&token, "id")
	//err := o.Read(&token)
	//if err != nil{
	//	fmt.Println("从数据查询失败!")
	//	return err
	//}
	accessToken, err := util.GetAccessToken(beego.AppConfig.String("appid"),beego.AppConfig.String("appsecret"),"https://api.weixin.qq.com/cgi-bin/token")
	if err != nil{
		fmt.Println("发送请求token失败!")
		return err
	}
	token.AccessToken = accessToken
	o.Update(&token,"access_token")
	return nil
}

func GetAccessToken(id int)(token *AccessToken, err error)  {
	o := orm.NewOrm()
	token.Id = id
	o.ReadOrCreate(&token,"id")
	return
}