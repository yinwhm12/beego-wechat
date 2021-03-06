package main

import (
	_ "beego-wechat/routers"

	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)



func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init()  {
	link := fmt.Sprintf("%s:%s@(%s:%s)/%s", beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqlport"), beego.AppConfig.String("mysqldb"))
	orm.RegisterDataBase("default", "mysql", link)

	orm.Debug = beego.BConfig.RunMode == "dev"
	//fmt.Println("appid=",beego.AppConfig.String("appid"),"appsecret=",beego.AppConfig.String("appsecret"))
	//go func() {
	//	time.Sleep(10 * time.Second)
	//	tool.ShopTimeTask()
	//}()
}