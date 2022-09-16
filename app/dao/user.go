/*
@Time : 2022/8/29 14:52
@Author : fushisanlang
@File : user
@Software: GoLand
*/
package dao

import (
	"farm/app/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func VerifyUserExist(userName string) bool {
	var (
		ctx = gctx.New()
		key = userName
	)
	UserExistStatus, err := g.Redis("data").Do(ctx, "exists", key)
	if err != nil {
		panic(err)
	}
	if gconv.Int(UserExistStatus) == 0 {
		return false
	} else {
		return true
	}

}
func RegisterUser(UserName string, RegisterUser g.Map) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "SET", UserName, RegisterUser)
	if err != nil {
		panic(err)
	}
}
func VerifyUser(userName string) *model.UserRegisterStruct {

	var (
		ctx = gctx.New()
		key = userName
	)
	result, err := g.Redis("data").Do(ctx, "GET", key)

	if err != nil {
		panic(err)
	}

	var UserStruct *model.UserRegisterStruct
	if err = result.Struct(&UserStruct); err != nil {
		panic(err)
	}
	return UserStruct
}
func GetUid(userName string) *model.UserRegisterStruct {

	var (
		ctx = gctx.New()
		key = userName
	)
	result, err := g.Redis("data").Do(ctx, "GET", key)

	if err != nil {
		panic(err)
	}
	var UserStruct *model.UserRegisterStruct
	if err = result.Struct(&UserStruct); err != nil {
		panic(err)
	}
	return UserStruct
}
