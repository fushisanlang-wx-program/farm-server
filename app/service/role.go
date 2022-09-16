/*
@Time : 2022/8/31 09:46
@Author : fushisanlang
@File : role
@Software: GoLand
*/
package service

import (
	"farm/app/dao"
	"farm/app/model"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetRoleByUid(uid string) (bool, int) {
	existsBool, roleCreateDone := dao.GetRoleByUid(uid)
	return existsBool, roleCreateDone
}
func CreateRoleName(uid, roleName string) {
	//fmt.Println("serviceid" + uid)

	roleStruct := model.RoleStruct{
		RoleName: roleName,
	}
	fmt.Println("26行" + roleName)
	dao.CreateRole(uid, roleStruct)
}

func CreateRoleStauts(uid string, jin, mu, shui, huo, tu int) bool {

	roleName := gconv.String(dao.GetRoleInfoByUid(uid, "RoleName").Array()[0])
	fmt.Println("33行" + roleName)

	if roleName != "" {
		roleStruct := model.RoleStruct{
			RoleName: roleName,
			Jin:      jin,
			Mu:       mu,
			Shui:     shui,
			Huo:      huo,
			Tu:       tu,
		}
		fmt.Println(roleStruct)
		dao.CreateRole(uid, roleStruct)
		return true
	} else {

		return false
	}
}

func CreateRoleTalented(uid string, talentedId int) bool {
	roleName := gconv.String(dao.GetRoleInfoByUid(uid, "RoleName").Array()[0])
	fmt.Println("56行" + roleName)

	if roleName != "" {

		jin := gconv.Int(dao.GetRoleInfoByUid(uid, "Jin").Array()[0])
		mu := gconv.Int(dao.GetRoleInfoByUid(uid, "Mu").Array()[0])
		shui := gconv.Int(dao.GetRoleInfoByUid(uid, "Shui").Array()[0])
		huo := gconv.Int(dao.GetRoleInfoByUid(uid, "Huo").Array()[0])
		tu := gconv.Int(dao.GetRoleInfoByUid(uid, "Tu").Array()[0])
		roleStruct := model.RoleStruct{
			RoleName:   roleName,
			Jin:        jin,
			Mu:         mu,
			Shui:       shui,
			Huo:        huo,
			Tu:         tu,
			TalentedId: talentedId,
		}
		dao.CreateRole(uid, roleStruct)
		return true
	} else {

		return false
	}
}
