/*
@Time : 2022/8/31 09:46
@Author : fushisanlang
@File : role
@Software: GoLand
*/
package dao

import (
	"farm/app/model"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

func GetRoleInfoByUid(uid, infoKey string) *gvar.Var {
	var ctx = gctx.New()
	v, err := g.Redis("data").Do(ctx, "HMGET", uid, infoKey)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	return v
}

func GetRoleByUid(uid string) (bool, int) {
	var (
		ctx        = gctx.New()
		key        = uid
		existsBool bool
	)
	var roleCreateDone int
	UserExistStatus, err := g.Redis("data").Do(ctx, "exists", key)
	if err != nil {
		panic(err)
	}
	if gconv.Int(UserExistStatus) == 0 {
		existsBool = false

	} else {
		existsBool = true

		v, err := g.Redis("data").Do(ctx, "HMGET", key, "RoleName")
		roleCreateDone = gconv.Int(v.Array()[0])

		if err != nil {
			panic(err)
		}
		DelRoleAtCreateErr(key)
	}

	return existsBool, roleCreateDone
}
func DelRoleAtCreateErr(uid string) {
	var ctx = gctx.New()
	fmt.Println("del" + uid)
	_, err := g.Redis("data").Do(ctx, "del", uid)
	if err != nil {
		panic(err)
	}

}
func CreateRole(uid string, roleStruct model.RoleStruct) {
	var (
		ctx = gctx.New()
	)

	_, err := g.Redis("data").Do(ctx, "HMSET", append(g.Slice{uid}, gutil.StructToSlice(roleStruct)...)...)

	if err != nil {
		panic(err)
	}

}
