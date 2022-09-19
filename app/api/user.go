package api

import (
	"farm/app/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.Request) {
	UserName := r.Get("UserName").String()
	Password := r.Get("Password").String()
	UserMail := r.Get("UserMail").String()

	if UserName == "" || Password == "" || UserMail == "" {
		ReturnErrCode(r, 417, "用户登录失败，数据空")
	} else if service.VerifyUserExist(UserName) == false {
		service.RegisterUser(UserName, UserMail, Password)
		r.Response.WriteJson(g.Map{
			"Message": "用户注册成功",
			"name":    UserName,
			"email":   UserMail,
			"code":    200,
		})

	} else {
		ReturnErrCode(r, 423, "用户注册失败，用户名冲突")

	}
}

func SignIn(r *ghttp.Request) {
	UserName := r.Get("UserName").String()
	Password := r.Get("Password").String()
	if UserName == "" || Password == "" {
		ReturnErrCode(r, 417, "用户登录失败，数据空")
	} else if service.VerifyUser(UserName, Password) == true {
		Uid := service.GetUid(UserName)
		r.Session.Set("UserName", UserName)
		r.Session.Set("Uid", Uid)
		r.Response.WriteJson(g.Map{
			"Message":  "用户登录成功",
			"UserName": UserName,
			"Uid":      Uid,
			"code":     200,
		})
	} else {
		ReturnErrCode(r, 401, "用户登录失败,账户密码不匹配")

	}
}
