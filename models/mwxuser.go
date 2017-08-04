package models

import "github.com/astaxie/beego/orm"

type WXUser struct {
	Id int	`json:"id,omitempty" orm:"pk;column(id);auto"`
	Subscribe	int	`json:"subscribe,omitempty" orm:"column(subscribe)"`
	Openid	string	`json:"openid,omitempty" orm:"column(openid)"`
	Nickname	string	`json:"nickname,omitempty" orm:"column(nickname)"`
	Sex int	`json:"sex,omitempty" orm:"column(sex)"`
	Language string `json:"language,omitempty" orm:"column(language)"`
	City	string	`json:"city,omitempty" orm:"column(city)"`
	Province string	`json:"province,omitempty" orm:"column(province)"`
	Country	string	`json:"country,omitempty" orm:"column(country)"`
	Headimgurl	string	`json:"headimgurl,omitempty" orm:"column(headimgurl)"`
	SubscribeTime int	`json:"subscribe_time,omitempty" orm:"column(subscribe_time)"`
	Unionid	string    `json:"unionid" orm:"column(unionid)"`
	Remark	string	`json:"remark,omitempty" orm:"column(remark)"`
	Groupid	int	`json:"groupid,omitempty" orm:"column(groupid)"`
	//TagidList	[]int	`json:"tagid_list,omitempty" orm:"column(tagid_list)"`
}

func (wxuser *WXUser)TableName() string  {
	return "wxuser"
}

func init()  {
	orm.RegisterModel(new(WXUser))
}

