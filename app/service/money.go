package service

import "farm/app/dao"

func getMoney(userName string) int {
	money := dao.GetMoney(userName)
	return money
}

func ChangeMoney(userName string, money int) {
	dao.ChangeMoney(userName, money)
	return
}
