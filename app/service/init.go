package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Return401(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"Message": "用户登录失败,账户密码不匹配",
		"code":    401,
	})
}

//func Return403(r *ghttp.Request) {
//	r.Response.WriteTpl("403.html", g.Map{
//		"gamename": GameName,
//	})
//}
