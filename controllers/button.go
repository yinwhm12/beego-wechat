package controllers

import "io/ioutil"

type WXMenuEvent struct {
	ToUserName	string `xml:"ToUserName"`
	FromUserName	string	`xml:"FromUserName"`
	CreateTime	int64	`xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Event	string `xml:"event"`
	EventKey	string	`xml:"event_key"`
	MenuId	string	`xml:"menu_id"`
	Content	string	`xml:"content"`
	Ticket string `xml:"ticket"`
}

type ButtonControllers struct {
	BaseControllers
}

// @router [post]
func (c *ButtonControllers)Post()  {
	if bytes, err := ioutil.ReadAll(c.Ctx.Request.Body); err == nil{

	}
}
