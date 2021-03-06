package models

import "github.com/astaxie/beego/orm"

type WXUser struct {
	Id int	`json:"id,omitempty" orm:"pk;column(id);auto"`
	//Subscribe	int	`json:"subscribe,omitempty" orm:"column(subscribe)"`
	Openid	string	`json:"openid,omitempty" orm:"column(openid)"`
	Nickname	string	`json:"nickname,omitempty" orm:"column(nickname)"`
	Sex int	`json:"sex,omitempty" orm:"column(sex)"`
	//Language string `json:"language,omitempty" orm:"column(language)"`
	Province string	`json:"province,omitempty" orm:"column(province)"`
	City	string	`json:"city,omitempty" orm:"column(city)"`
	Country	string	`json:"country,omitempty" orm:"column(country)"`
	Headimgurl	string	`json:"headimgurl,omitempty" orm:"column(headimgurl)"`
	//SubscribeTime int	`json:"subscribe_time,omitempty" orm:"column(subscribe_time)"`
	Unionid	string    `json:"unionid" orm:"column(unionid)"`
	Privilege	string	`json:"privilege,omitempty" orm:"column(privilege)"`
	//Remark	string	`json:"remark,omitempty" orm:"column(remark)"`
	//Groupid	int	`json:"groupid,omitempty" orm:"column(groupid)"`
	//TagidList	[]int	`json:"tagid_list,omitempty" orm:"column(tagid_list)"`
}

func (wxuser *WXUser)TableName() string  {
	return "wxuser"
}

func init()  {
	orm.RegisterModel(new(WXUser))
}

func AddWXUser(wxuser WXUser)(error)  {
	o := orm.NewOrm()
	_, err := o.Insert(&wxuser)
	return err
}

func GetWXUserByOpenid(openid string) (wxuser WXUser, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(ARToken)).Filter("Openid",openid).One(&wxuser)
	return
}

func GetWXUserById(id int)(wxuser WXUser,err error)  {
	o := orm.NewOrm()
	wxuser = WXUser{Id:id}
	if err = o.Read(&wxuser); err == nil{
		return wxuser, nil
	}
	return wxuser, err
}
