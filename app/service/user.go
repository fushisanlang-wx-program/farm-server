/*
@Time : 2022/8/29 14:50
@Author : fushisanlang
@File : user
@Software: GoLand
*/
package service

import (
	"farm/app/dao"
	"farm/app/model"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

func VerifyUserExist(userName string) bool {
	UserExistStatus := dao.VerifyUserExist(userName)
	return UserExistStatus

}
func VerifySession(r *ghttp.Request) (bool, string, string) {
	sessionData, err := r.Session.Data()
	if err != nil {
		Return401(r)
		return false, "", ""
	}
	var UserStruct *model.UserRegisterStruct
	if gconv.Struct(sessionData, &UserStruct) != nil {
		Return401(r)
		return false, "", ""
	}
	userName := UserStruct.UserName
	uid := UserStruct.Uid

	uidTrue := GetUid(userName)
	if uid == uidTrue {
		return true, uidTrue, userName
	} else {
		Return401(r)
		return false, "", ""
	}

}
func RegisterUser(UserName, UserMail, Password string) {

	RegisterUser := g.Map{
		"UserName":   UserName,
		"Password":   Password,
		"UserMail":   UserMail,
		"ReliveTime": 0,
		"Uid":        guid.S(),
	}

	dao.RegisterUser(UserName, RegisterUser)
}
func VerifyUser(userName, password string) bool {
	if VerifyUserExist(userName) == true {
		passwordTrue := dao.VerifyUser(userName).Password
		if password == passwordTrue {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func GetUid(userName string) string {
	userStruct := dao.GetUid(userName)
	if userStruct == nil {
		return ""
	} else {
		Uid := userStruct.Uid
		return Uid
	}
}
func GetUserInfo(userName string) *model.UserRegisterStruct {
	userStruct := dao.GetUid(userName)
	return userStruct
}

/*
//获取session中的username
func GetSessionUserName(r *ghttp.Request) (bool, interface{}) {
	a, _ := r.Session.Data()["UserName"]
	//b := fmt.Sprintf("%v", a)
	//fmt.Println(b)
	if a == nil {

		return false, ""
	} else {

		return true, a
	}
}

//获取session中的uid
func GetSessionUserId(r *ghttp.Request) (bool, int) {

	a := r.Session.Map()["Uid"]

	if a == nil {

		return false, 0
	} else {
		uid64, _ := a.(json.Number).Int64()

		uid := *(*int)(unsafe.Pointer(&uid64))
		return true, uid
	}

}
*/
