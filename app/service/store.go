package service

import (
	"farm/app/dao"
	"farm/app/model"
)

var StoreList [50]model.PlantInfoStruct

func StoreBuy(uId, userName string, plantId, buyCount int) (bool, string) {
	plantInfo := dao.GetPlantInfo(plantId)
	userLevel := GetLevel(getEx(userName))
	userMoney := getMoney(userName)
	needMoney := plantInfo.BuyPrice * buyCount
	if userLevel >= plantInfo.Level {
		if userMoney >= needMoney {
			ChangeMoney(userName, userMoney-needMoney)
			plantCount := getPlantCount(uId, plantId)
			SetPlantCount(uId, plantId, plantCount+buyCount)
			return true, ""
		} else {
			return false, "剩余钱不足"
		}
	} else {
		return false, "等级不足"
	}

}

func init() {
	StoreList = dao.GetStoreList()
}
