/*
@Time : 2022/8/29 10:37
@Author : fushisanlang
@File : textscript
@Software: GoLand
*/
package dao

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetTextScript(textName string) string {
	var (
		ctx = gctx.New()
		key = textName
	)
	result, err := g.Redis("config").Do(ctx, "GET", key)
	if err != nil {
		panic(err)
	}
	return gconv.String(result)

}
