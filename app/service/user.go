package service

import (
	"farm/app/dao"
	"farm/app/model"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/util/guid"
)

func VerifyUserExist(userName string) bool {
	userExistStatus := dao.VerifyUserExist(userName)
	return userExistStatus

}
func VerifySession(r *ghttp.Request) (bool, string, string) {
	sessionData, err := r.Session.Data()
	if err != nil {
		return false, "", ""
	}
	var userStruct *model.UserRegisterStruct

	if gconv.Struct(sessionData, &userStruct) != nil {
		return false, "", ""
	}
	userName := userStruct.UserName
	if VerifyUserExist(userName) == true {
		uid := userStruct.Uid

		uidTrue := GetUid(userName)
		if uid == uidTrue {
			return true, uidTrue, userName
		} else {
			return false, "", ""
		}
	} else {
		return false, "", ""
	}

}
func RegisterUser(userName, userMail, password string) {
	UId := guid.S()

	RegisterUser := model.UserRegisterStruct{
		UserName:   userName,
		Password:   password,
		EMail:      userMail,
		ReliveTime: 0,
		Uid:        UId,
		FieldCount: 6,
		Money:      2000,
	}

	dao.RegisterUser(userName, RegisterUser)
	registerFarm(UId)

	for i := 1; i < 7; i++ {
		dao.OpenField(userName, UId, i)
	}
	registerBag(UId)
}
func VerifyUser(userName, password string) bool {
	if VerifyUserExist(userName) == true {
		passwordTrue := dao.GetUserPass(userName)
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
