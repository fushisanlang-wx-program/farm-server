package dao

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetMoney(UserName string) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("data").Do(ctx, "HMGET", UserName, "Money")
	if err != nil {
		panic(err)
	}
	return result.Int()

}

func GetOpenFieldNeedMoney(fieldId int) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("config").Do(ctx, "HMGET", "field"+gconv.String(fieldId), "Money")
	if err != nil {
		panic(err)
	}
	LevelNeedEx := result.Int()
	return LevelNeedEx

}

func ChangeMoney(UserName string, Money int) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", UserName, "Money", Money)
	if err != nil {
		panic(err)
	}
}
