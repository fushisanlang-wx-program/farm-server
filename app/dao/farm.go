package dao

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func RegisterFarm(Uid int) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "SET", Uid, RegisterUser)
	if err != nil {
		panic(err)
	}
}
