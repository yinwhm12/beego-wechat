package tool

import (
	"beego-wechat/util"
	"strings"
)

//处理用户点击菜单 或者关注以及取消
func DealMenuEvent(ne *util.NormalEvent)  {
	switch ne.Event {
	case "VIEW":
	case "CLICK":
		if strings.EqualFold(ne.EventKey,""){

		}else if strings.EqualFold(ne.EventKey,""){

		}else{

		}
	case "subscribe":
		if strings.Contains(ne.EventKey,"qrscene_"){//扫码事件

		}else{

		}
	case "unsubscribe":

	case  "LOCATION"://上报地理位置事件
	case "SCAN"://已关注的 扫他人二维码

	default:

	}
}