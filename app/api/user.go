/*
@Time : 2022/8/26 10:35
@Author : fushisanlang
@File : status
@Software: GoLand
*/
package api

import (
	"farm/app/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

/*
//世界状态
func GetStatus(r *ghttp.Request) {
	WorldMp, WorldGeneration := "1","2"
	r.Response.WriteJson(g.Map{
		"WorldMp":    WorldMp,
		"Generation": WorldGeneration,
	})
}
*/
//服务器版本
func SignIn(r *ghttp.Request) {
	r.Response.WriteTpl("signin.html")

}
func Register(r *ghttp.Request) {
	r.Response.WriteTpl("register.html")

}

func RegisterVerify(r *ghttp.Request) {
	UserName := r.GetForm("username").String()
	UserMail := r.GetForm("useradress").String()
	Password := r.GetForm("password").String()
	if service.VerifyUserExist(UserName) == false {
		service.RegisterUser(UserName, UserMail, Password)
		r.Response.RedirectTo("/user/signin")
	} else {
		r.Response.WriteTpl("registeragain.html")
	}
}

func VerifyUser(r *ghttp.Request) {
	UserName := r.GetForm("username").String()
	Password := r.GetForm("password").String()
	if service.VerifyUser(UserName, Password) == true {
		Uid := service.GetUid(UserName)
		r.Session.Set("UserName", UserName)
		r.Session.Set("Uid", Uid)
		r.Response.RedirectTo("/role/info")
	} else {
		r.Response.WriteTpl("signinagain.html")
	}
}
