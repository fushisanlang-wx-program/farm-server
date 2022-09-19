package service

import (
	"farm/app/dao"
	"farm/app/model"
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

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
	UId := guid.S()

	RegisterUser := model.UserRegisterStruct{
		UserName:   UserName,
		Password:   Password,
		EMail:      UserMail,
		ReliveTime: 0,
		Uid:        UId,
		FieldCount: 6,
		Money:      2000,
	}

	dao.RegisterUser(UserName, RegisterUser)
	registerFarm(UserName, UId)
	for i := 1; i < 7; i++ {
		OpenField(UserName, UId)
	}
}
func VerifyUser(userName, password string) bool {
	if VerifyUserExist(userName) == true {
		passwordTrue := dao.GetUserPass(userName)
		fmt.Println(passwordTrue)
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
	if VerifyUserExist(userName) == true {
		Uid := dao.GetUid(userName)
		return Uid
	} else {
		return ""
	}
}

//func GetUserInfo(userName string) *model.UserRegisterStruct {
//	userStruct := dao.GetUid(userName)
//	return userStruct
//}
//
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
