package service

import "farm/app/dao"

func GetLevel(ex int) int {
	level := 0
	switch {
	case ex < 200:
		level = 0
	case ex >= 200 && ex < 600:
		level = 1
	case ex >= 600 && ex < 1200:
		level = 2
	case ex >= 1200 && ex < 2000:
		level = 3
	case ex >= 2000 && ex < 3000:
		level = 4
	case ex >= 3000 && ex < 4200:
		level = 5
	case ex >= 4200 && ex < 5600:
		level = 6
	case ex >= 5600 && ex < 7200:
		level = 7
	case ex >= 7200 && ex < 9000:
		level = 8
	case ex >= 9000 && ex < 11000:
		level = 9
	case ex >= 11000 && ex < 13200:
		level = 10
	case ex >= 13200 && ex < 15600:
		level = 11
	case ex >= 15600 && ex < 18200:
		level = 12
	case ex >= 18200 && ex < 21000:
		level = 13
	case ex >= 21000 && ex < 24000:
		level = 14
	case ex >= 24000:
		level = 15

	}
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
