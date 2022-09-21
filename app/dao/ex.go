package dao

import (
	"farm/app/logger"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetEx(userName string) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("data").Do(ctx, "HMGET", userName, "Ex")
	if err != nil {
		logger.LogError("获取用户经验失败，用户名：" + userName)
		panic(err)
	}
	return gconv.Int(result.Array()[0])

}
func GetOpenFieldNeedLevel(fieldId int) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("config").Do(ctx, "HMGET", "field_"+gconv.String(fieldId), "Level")
	if err != nil {
		logger.LogError("获取土地等级信息失败，土地ID：" + gconv.String(fieldId))
		panic(err)
	}
	OpenFieldNeedLevel := gconv.Int(result.Array()[0])
	return OpenFieldNeedLevel

}
func ChangeEx(userName string, Ex int) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", userName, "Ex", Ex)
	if err != nil {
		logger.LogError("修改用户经验信息失败，用户名：" + userName + "，经验值：" + gconv.String(Ex))

		panic(err)
	}
}
