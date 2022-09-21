package logger

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func LogInfo(logString string) {
	ctx := context.TODO()
	g.Log().Info(ctx, logString)
}
func LogError(logString string) {
	ctx := context.TODO()
	g.Log().Error(ctx, logString)
}
