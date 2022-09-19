package dao

import (
	"farm/app/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
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
func RegisterUser(UserName string, RegisterUser model.UserRegisterStruct) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", append(g.Slice{UserName}, gutil.StructToSlice(RegisterUser)...)...)
	if err != nil {
		panic(err)
	}
}
func GetUserPass(userName string) string {

	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("data").Do(ctx, "HMGET", userName, "Password")
	if err != nil {
		panic(err)
	}

	return gconv.String(result.Array()[0])
}
func GetUid(userName string) string {

	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("data").Do(ctx, "HMGET", userName, "Uid")

	if err != nil {
		panic(err)
	}

	return gconv.String(result.Array()[0])
}

func GetUserInfoFieldCount(userName string) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("data").Do(ctx, "HMGET", userName, "FieldCount")
	if err != nil {
		panic(err)
	}

	return gconv.Int(result.Array()[0])
}
