/*
@Time : 2022/8/29 16:59
@Author : fushisanlang
@File : role
@Software: GoLand
*/
package api

import (
	"farm/app/service"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

func RoleInfo(r *ghttp.Request) {
	verifyStatus, uid, _ := service.VerifySession(r)

	if verifyStatus == true {
		existsBool, roleStruct := service.GetRoleByUid(uid)
		//fmt.Println(existsBool, roleStruct)
		if existsBool == true && roleStruct == 1 {

			r.Response.Write(roleStruct)

		} else {

			r.Response.RedirectTo("/role/create/page1")
		}
	}
}
func RoleCreatePage1(r *ghttp.Request) {
	verifyStatus, uid, userName := service.VerifySession(r)

	if verifyStatus == true {
		userReliveTime := service.GetUserInfo(userName).ReliveTime

		GameName := service.GameName
		r.Response.WriteTpl("/role/create/page1.html", g.Map{
			"gamename":    GameName,
			"createinfo":  "请输入道友第" + gconv.String(userReliveTime+1) + "次转生信息" + uid,
			"createalert": "",
		})
	}
}

func RoleCreatePage1Verify(r *ghttp.Request) {
	//fmt.Println(1)
	verifyStatus, uid, userName := service.VerifySession(r)
	userReliveTime := service.GetUserInfo(userName).ReliveTime
	GameName := service.GameName
	if verifyStatus == true {
		rolename := r.GetForm("rolename").String()
		if rolename == "" {
			r.Response.WriteTpl("/role/create/page1.html", g.Map{
				"gamename":    GameName,
				"createinfo":  "请输入道友第" + gconv.String(userReliveTime+1) + "次转生信息" + uid,
				"createalert": "道号不能为空",
			})
		} else {
			fmt.Println(rolename)
			service.CreateRoleName(uid, rolename)
			//fmt.Println("apiid" + uid)
			r.Response.RedirectTo("/role/create/page2")
		}

	}

}
func RoleCreatePage2(r *ghttp.Request) {
	verifyStatus, _, _ := service.VerifySession(r)

	if verifyStatus == true {
		//userReliveTime := service.GetUserInfo(userName).ReliveTime

		GameName := service.GameName
		r.Response.WriteTpl("/role/create/page2.html", g.Map{
			"gamename":    GameName,
			"createinfo":  "", // "请输入道友第" + gconv.String(userReliveTime+1) + "次转生信息" + uid,
			"createalert": "",
		})
	}
}

func RoleCreatePage2Verify(r *ghttp.Request) {
	verifyStatus, uid, _ := service.VerifySession(r)
	GameName := service.GameName
	if verifyStatus == true {
		statusStr := r.GetForm("status").String()
		statusArr := strings.Split(statusStr, ",")
		if len(statusArr) == 5 {
			jin := gconv.Int(statusArr[0])
			mu := gconv.Int(statusArr[1])
			shui := gconv.Int(statusArr[2])
			huo := gconv.Int(statusArr[3])
			tu := gconv.Int(statusArr[4])
			if jin+mu+shui+huo+tu == 20 {
				if service.CreateRoleStauts(uid, jin, mu, shui, huo, tu) == true {

					//fmt.Println("apiid" + uid)
					//r.Response.RedirectTo("/role/create/page2")
					r.Response.RedirectTo("/role/create/page3")
				} else {
					r.Response.RedirectTo("/role/create/page1")
				}
			} else {
				r.Response.WriteTpl("/role/create/page2.html", g.Map{
					"gamename":    GameName,
					"createinfo":  "", // "请输入道友第" + gconv.String(userReliveTime+1) + "次转生信息" + uid,
					"createalert": "属性异常，请重新选择属性",
				})
			}
		} else {
			r.Response.WriteTpl("/role/create/page2.html", g.Map{
				"gamename":    GameName,
				"createinfo":  "", // "请输入道友第" + gconv.String(userReliveTime+1) + "次转生信息" + uid,
				"createalert": "属性异常，请重新选择属性",
			})
		}
	}

}

func RoleCreatePage3(r *ghttp.Request) {
	verifyStatus, _, _ := service.VerifySession(r)

	if verifyStatus == true {
		//userReliveTime := service.GetUserInfo(userName).ReliveTime

		GameName := service.GameName
		r.Response.WriteTpl("/role/create/page3.html", g.Map{
			"gamename":    GameName,
			"createinfo":  "", // "请输入道友第" + gconv.String(userReliveTime+1) + "次转生信息" + uid,
			"createalert": "",
		})
	}
}

func RoleCreatePage3Verify(r *ghttp.Request) {
	verifyStatus, uid, _ := service.VerifySession(r)
	GameName := service.GameName
	if verifyStatus == true {

		statusStr := r.GetForm("status").String()

		statusArr := strings.Split(statusStr, ",")
		if len(statusArr) == 2 {
			aId := gconv.Int(statusArr[0])
			bId := gconv.Int(statusArr[1])
			if aId >= 1 && aId <= 5 && bId <= 6 && bId >= 1 {
				talentedId := (aId-1)*6 + bId
				if service.CreateRoleTalented(uid, talentedId) == true {
					r.Response.Write(talentedId)
				} else {
					r.Response.RedirectTo("/role/create/page1")

				}
			} else {
				r.Response.WriteTpl("/role/create/page3.html", g.Map{
					"gamename":    GameName,
					"createinfo":  "", // "请输入道友第" + gconv.String(userReliveTime+1) + "次转生信息" + uid,
					"createalert": "属性异常，请重新选择属性",
				})
			}
		} else {
			r.Response.WriteTpl("/role/create/page3.html", g.Map{
				"gamename":    GameName,
				"createinfo":  "", // "请输入道友第" + gconv.String(userReliveTime+1) + "次转生信息" + uid,
				"createalert": "属性异常，请重新选择属性",
			})
		}
	}

}
