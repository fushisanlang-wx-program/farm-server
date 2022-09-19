package service

import "farm/app/dao"

func registerFarm(uid int) {
	dao.RegisterFarm(uid)
}
