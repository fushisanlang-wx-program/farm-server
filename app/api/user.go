package api

import (
	"farm/app/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func UserRegister(r *ghttp.Request) {
	userName := r.Get("UserName").String()
	password := r.Get("Password").String()
	userMail := r.Get("UserMail").String()

	if userName == "" || password == "" || userMail == "" {
		returnErrCode(r, 417, "用户登录失败，数据空")
	} else if service.VerifyUserExist(userName) == false {
		service.RegisterUser(userName, userMail, password)
		r.Response.WriteJson(g.Map{
			"message": "用户注册成功",
			"name":    userName,
			"email":   userMail,
		})

	} else {
		returnErrCode(r, 423, "用户注册失败，用户名已存在")

	}
}

func UserSignIn(r *ghttp.Request) {
	userName := r.Get("UserName").String()
	password := r.Get("Password").String()
	if userName == "" || password == "" {
		returnErrCode(r, 417, "用户登录失败，数据空")
	} else if service.VerifyUser(userName, password) == true {
		Uid := service.GetUid(userName)
		r.Session.Set("userName", userName)
		r.Session.Set("Uid", Uid)
		r.Response.WriteJson(g.Map{
			"message":  "用户登录成功",
			"userName": userName,
		})
	} else {
		returnErrCode(r, 401, "用户登录失败,账户密码不匹配")

	}
}
