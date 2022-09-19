package dao

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetEx(UserName string) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("data").Do(ctx, "HMGET", UserName, "Ex")
	if err != nil {
		panic(err)
	}
	return result.Int()

}
func GetOpenFieldNeedLevel(fieldId int) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("config").Do(ctx, "HMGET", "field"+gconv.String(fieldId), "Level")
	if err != nil {
		panic(err)
	}
	LevelNeedEx := result.Int()
	return LevelNeedEx

}
func ChangeEx(UserName string, Ex int) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", UserName, "Ex", Ex)
	if err != nil {
		panic(err)
	}
}
