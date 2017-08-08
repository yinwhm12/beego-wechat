package util


//订阅 以及取消订阅
type SubscribeEvent struct {
	ToUserName	string `xml:"ToUserName"`
	FromUserName	string `xml:"FromUserName"`
	CreateTime	int64	`xml:"CreateTime"`
	MsgType	string `xml:"MsgType"`
	Event string `xml:"Event"`
}

//用户已关注时的事件推送
type NormalEvent struct {
	ToUserName	string `xml:"ToUserName"`
	FromUserName	string `xml:"FromUserName"`
	CreateTime	int64	`xml:"CreateTime"`
	MsgType	string `xml:"MsgType"`
	Event string `xml:"Event"`
	EventKey string `xml:"EventKey"`
	Ticket	string 	`xml:"Ticket"`
}