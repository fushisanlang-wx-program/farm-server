package dao

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetMoney(userName string) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("data").Do(ctx, "HMGET", userName, "Money")
	if err != nil {
		panic(err)
	}
	return gconv.Int(result.Array()[0])

}

func GetOpenFieldNeedMoney(fieldId int) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("config").Do(ctx, "HMGET", "field_"+gconv.String(fieldId), "Money")
	if err != nil {
		panic(err)
	}
	OpenFieldNeedMoney := gconv.Int(result.Array()[0])
	return OpenFieldNeedMoney

}

func ChangeMoney(userName string, money int) {
	var (
		ctx = gctx.New()
	)
	fmt.Println(userName)
	fmt.Println(money)
	_, err := g.Redis("data").Do(ctx, "HMSET", userName, "Money", money)
	if err != nil {
		panic(err)
	}
}
