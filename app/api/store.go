package api

import (
	"farm/app/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func StoreBuy(r *ghttp.Request) {
	plantId := r.Get("plantId").Int()
	buyCount := r.Get("buyCount").Int()
	if plantId == 0 || buyCount == 0 {
		returnErrCode(r, 417, "数据空")
	} else {
		verifyStatus, uId, userName := service.VerifySession(r)
		if verifyStatus == true {
			buyStatus, errMsg := service.StoreBuy(uId, userName, plantId, buyCount)

			if buyStatus == true {
				r.Response.WriteJson(g.Map{
					"plantId":  plantId,
					"buyCount": buyCount,
				})
			} else {
				returnErrCode(r, 417, errMsg)
			}

		} else {
			returnErrCode(r, 401, "用户校验失败,请重新登录")
		}
	}
}
func StoreList(r *ghttp.Request) {

	verifyStatus, _, _ := service.VerifySession(r)
	if verifyStatus == true {

		r.Response.WriteJson(service.StoreList)

	} else {
		returnErrCode(r, 401, "用户校验失败,请重新登录")
	}
}
