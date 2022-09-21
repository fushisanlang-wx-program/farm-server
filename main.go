package main

import (
	"farm/app"
	"farm/app/logger"
)

func main() {
	logger.LogInfo("服务启动")
	app.Run()

}
