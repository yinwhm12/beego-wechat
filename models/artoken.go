package models

import "github.com/astaxie/beego/orm"

type ARToken struct {
	Id int	`json:"id,omitempty" orm:"pk;column(id);auto"`
	AccessToken	string	`json:"access_token,omitempty" orm:"column(access_token);size(256)"`
	ExpiresIn int64 `json:"expires_in,omitempty" orm:"column(expires_in)"`
	RefreshToken	string	`json:"refresh_token,omitempty" orm:"column(refresh_token);size(256)"`
	Openid	string	`json:"openid,omitempty" orm:"column(openid)"`
	Scope string	`json:"scope,omitempty" orm:"column(scope)"`
}

func (this *ARToken)TableName()string  {
	return "artoken"
}

func init()  {
	orm.RegisterModel(new(ARToken))
}

//func GetAndUpdateARTokenByOpenid(openid string)(err error)  {
//	o := orm.NewOrm()
//	artoken := ARToken{Openid:openid}
//	o.ReadOrCreate(&artoken,"openid")
//	o.Update(artoken,"")
//}

func GetARTokenByOpenid(openid string)(art ARToken, err error)  {
	o := orm.NewOrm()
	err = o.QueryTable(new(ARToken)).Filter("Openid",openid).One(&art)
	return
}

func UpdateARTokenByOpenid(art *ARToken)(err error)  {
	o := orm.NewOrm()
	_, err = o.Update(art, "AccessToken","RefreshToken")
	return
}

func GetARTokenById(id int) (art *ARToken, err error) {
	o := orm.NewOrm()
	//art := ARToken{Id:id}
	art = &ARToken{Id:id}
	if err = o.Read(art); err == nil{
		return art, nil
	}
	return nil, err
}

func AddARToken(art ARToken)(err error){
	o := orm.NewOrm()
	_, err = o.Insert(&art)
	return
}


