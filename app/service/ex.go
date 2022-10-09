package service

import "farm/app/dao"

func getLevel(ex int) (int, int) {
	level := 0
	needEx := 0
	for {
		needEx := (level + 1) * (level + 2) * 100
		if needEx > ex {

			break
		}
		level = level + 1
	}
	ex = ex - needEx
	return level, ex

}

func GetLevel(ex int) int {
	level := 0
	needEx := 0
	for {
		needEx := (level + 1) * (level + 2) * 100
		if needEx > ex {

			break
		}
		level = level + 1
	}
	ex = ex - needEx
	return level

}
func getEx(userName string) int {
	ex := dao.GetEx(userName)
	return ex
}
func ChangeEx(userName string, ex int) {
	dao.ChangeEx(userName, ex)
	//增加判断是否升级，
}
