package main

import (
	"context"
	"farm/app"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	ctx := context.TODO()
	g.Log().Info(ctx, "程序启动")
	app.Run()

}
