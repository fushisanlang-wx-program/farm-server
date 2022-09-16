/*
@Time : 2022/8/26 16:42
@Author : fushisanlang
@File : init
@Software: GoLand
*/
package service

import (
	"farm/app/dao"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

var GameName, Version string

func init() {

	setStatus()
}

func setStatus() {
	var ctx = gctx.New()
	VersionFromConf, _ := gcfg.Instance().Get(ctx, "status.version")
	GameNameFromConf, _ := gcfg.Instance().Get(ctx, "status.gamename")
	Status := g.Map{
		"GameName": GameNameFromConf,
		"Version":  VersionFromConf,
	}
	dao.SetStatus(Status)
	StatusStruct := dao.GetStatus()
	Version = StatusStruct.Version
	GameName = StatusStruct.GameName

}

func Return401(r *ghttp.Request) {
	//r.Response.RedirectTo("/user/signin")
	r.Response.WriteTpl("401.html", g.Map{
		"gamename": GameName,
	})
}

//func Return403(r *ghttp.Request) {
//	r.Response.WriteTpl("403.html", g.Map{
//		"gamename": GameName,
//	})
//}
