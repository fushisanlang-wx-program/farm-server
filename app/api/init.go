/*
@Time : 2022/8/29 14:03
@Author : fushisanlang
@File : init.go
@Software: GoLand
*/
package api

import (
	"farm/app/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func RootPage(r *ghttp.Request) {
	GameName := service.GameName
	//	r.Response.WriteJson(g.Map{
	//		"Message": "登录成功", "code": 200,
	//})
	r.Response.WriteTpl("index.html", g.Map{
		"gamename":  GameName,
		"indextext": service.GetTextScript("textScript_index"),

		"indexfrom": service.GetTextScript("textScript_indexfrom"),
	})

}
