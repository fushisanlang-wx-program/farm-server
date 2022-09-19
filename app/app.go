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
	s.SetSessionStorage(gsession.NewStorageRedis(g.Redis("session")))
	//s.BindHandler("/", api.RootPage)
	//group := s.Group("/status")
	//服务器版本
	//group.ALL("/version", api.GetVersion)
	group := s.Group("/user")
	group.POST("/register", api.Register)
	group.POST("/signin", api.SignIn)
	group = s.Group("/field")
	group.POST("/signin", api.SignIn)

	s.Run()
}
