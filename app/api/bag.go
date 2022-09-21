package api

import (
	"farm/app/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func BagInfo(r *ghttp.Request) {
	verifyStatus, uId, _ := service.VerifySession(r)
	if verifyStatus == true {
		bagInfo := service.GetBagInfo(uId)
		r.Response.WriteJson(bagInfo)
	} else {
		returnErrCode(r, 401, "用户校验失败,请重新登录")
	}
}
