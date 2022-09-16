/*
@Time : 2022/8/26 10:31
@Author : fushisanlang
@File : app.go
@Software: GoLand
*/
package app

import (
	"farm/app/api"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gsession"
	"time"
)

func Run() {
	//配置文件
	//g.Cfg().SetFileName("config.yml")
	//服务定义
	s := g.Server()
	//session相关

	s.SetSessionMaxAge(time.Hour)
	s.SetSessionStorage(gsession.NewStorageRedis(g.Redis()))

	s.BindHandler("/", api.RootPage)
	//group := s.Group("/status")
	//服务器版本
	//group.ALL("/version", api.GetVersion)
	group := s.Group("/user")
	group.ALL("/signin", api.SignIn)
	group.ALL("/verify", api.VerifyUser)
	group.ALL("/register", api.Register)
	group.ALL("/registerverify", api.RegisterVerify)
	group = s.Group("/role")
	group.ALL("/info", api.RoleInfo)
	group.ALL("/create/page1", api.RoleCreatePage1)
	group.ALL("/create/page1/verify", api.RoleCreatePage1Verify)
	group.ALL("/create/page2", api.RoleCreatePage2)
	group.ALL("/create/page2/verify", api.RoleCreatePage2Verify)
	group.ALL("/create/page3", api.RoleCreatePage3)
	group.ALL("/create/page3/verify", api.RoleCreatePage3Verify)
	//启动web
	s.Run()
}
