/*
@Time : 2022/8/26 10:51
@Author : fushisanlang
@File : status
@Software: GoLand
*/
package dao

import (
	"farm/app/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func GetStatus() *model.StatusStruct {
	var (
		ctx = gctx.New()
		key = "Status"
	)
	result, err := g.Redis("config").Do(ctx, "GET", key)
	if err != nil {
		panic(err)
	}

	var StatusStruct *model.StatusStruct
	if err = result.Struct(&StatusStruct); err != nil {
		panic(err)
	}
	return StatusStruct
}
func SetStatus(statusMap g.Map) {
	var (
		ctx = gctx.New()

		key = "Status"
	)
	_, err := g.Redis("config").Do(ctx, "SET", key, statusMap)
	if err != nil {
		panic(err)
	}

}
