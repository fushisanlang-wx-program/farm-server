package api

import (
	"farm/app/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func FieldOpen(r *ghttp.Request) {
	verifyStatus, uId, userName := service.VerifySession(r)
	if verifyStatus == true {
		status, fieldId := service.OpenField(userName, uId)
		if status == true {
			r.Response.WriteJson(g.Map{
				"fieldId": fieldId,
			})
		} else {
			returnErrCode(r, 423, "开启失败，资源不足")
		}
	} else {
		returnErrCode(r, 401, "用户校验失败,请重新登录")
	}
}

func FieldInfo(r *ghttp.Request) {
	verifyStatus, uId, _ := service.VerifySession(r)
	if verifyStatus == true {
		fieldInfo := service.GetFieldInfo(uId)
		r.Response.WriteJson(fieldInfo)
	} else {
		returnErrCode(r, 401, "用户校验失败,请重新登录")
	}
}
func FieldInfoById(r *ghttp.Request) {
	fieldId := r.Get("fieldId").Int()
	if fieldId > 18 || fieldId < 1 {
		returnErrCode(r, 404, "id异常")
	} else {
		verifyStatus, uId, _ := service.VerifySession(r)
		if verifyStatus == true {
			fieldInfo := service.FieldInfoById(uId, fieldId)
			r.Response.WriteJson(fieldInfo)
		} else {
			returnErrCode(r, 401, "用户校验失败,请重新登录")
		}
	}
}
func FieldPlant(r *ghttp.Request) {
	fieldId := r.Get("fieldId").Int()
	plantId := r.Get("plantId").Int()

	if fieldId > 18 || fieldId < 1 || plantId < 1 || plantId > 50 {
		returnErrCode(r, 404, "id异常")
	} else {
		verifyStatus, uId, userName := service.VerifySession(r)
		if verifyStatus == true {
			plantStatus := service.FieldPlant(uId, userName, fieldId, plantId)
			if plantStatus == true {
				fieldInfo := service.FieldInfoById(uId, fieldId)
				r.Response.WriteJson(fieldInfo)
			} else {
				returnErrCode(r, 423, "种植失败，资源异常")
			}
		} else {
			returnErrCode(r, 401, "用户校验失败,请重新登录")
		}
	}
}

func FieldHarvest(r *ghttp.Request) {
	fieldId := r.Get("fieldId").Int()
	if fieldId > 18 || fieldId < 1 {
		returnErrCode(r, 404, "id异常")
	} else {
		verifyStatus, uId, userName := service.VerifySession(r)
		if verifyStatus == true {

			harvestStatus, harvestInfo := service.FieldHarvest(uId, userName, fieldId)
			if harvestStatus == true {
				r.Response.WriteJson(harvestInfo)
			} else {
				returnErrCode(r, 423, "收获失败，资源异常")
			}

		} else {
			returnErrCode(r, 401, "用户校验失败,请重新登录")
		}
	}
}

func FieldEradicate(r *ghttp.Request) {
	fieldId := r.Get("fieldId").Int()
	if fieldId > 18 || fieldId < 1 {
		returnErrCode(r, 404, "id异常")
	} else {
		verifyStatus, uId, userName := service.VerifySession(r)
		if verifyStatus == true {

			eradicateStatus := service.FieldEradicate(uId, userName, fieldId)
			if eradicateStatus == true {
				fieldInfo := service.FieldInfoById(uId, fieldId)
				r.Response.WriteJson(fieldInfo)
			} else {
				returnErrCode(r, 423, "提升失败，资源异常")
			}

		} else {
			returnErrCode(r, 401, "用户校验失败,请重新登录")
		}
	}
}
func FieldUpgrade(r *ghttp.Request) {
	fieldId := r.Get("fieldId").Int()
	if fieldId > 18 || fieldId < 1 {
		returnErrCode(r, 404, "id异常")
	} else {
		verifyStatus, uId, userName := service.VerifySession(r)
		if verifyStatus == true {

			upgradeStatus := service.FieldUpgrade(uId, userName, fieldId)
			if upgradeStatus == true {
				fieldInfo := service.FieldInfoById(uId, fieldId)
				r.Response.WriteJson(fieldInfo)
			} else {
				returnErrCode(r, 423, "提升失败，资源异常")
			}

		} else {
			returnErrCode(r, 401, "用户校验失败,请重新登录")
		}
	}
}
