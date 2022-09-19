package service

import (
	"farm/app/dao"
	"farm/app/model"
)

func registerFarm(UserName, uId string) {
	for i := 1; i <= 18; i++ {
		FarmFieldInfoStruct := model.FarmFieldInfoStruct{
			UserId:         uId,
			FieldId:        i,
			Status:         0,
			PlantID:        0,
			MaturationTime: 0,
			ReMature:       0,
		}
		dao.RegisterFarm(uId, FarmFieldInfoStruct)

	}
}

func OpenField(userName, uId string) bool {
	userFieldCount := dao.GetUserInfoFieldCount(userName)
	if userFieldCount < 18 {
		newFieldId := userFieldCount + 1
		userEx := dao.GetEx(userName)
		userLevel := GetLevel(userEx)
		userMoney := dao.GetMoney(userName)
		openFieldNeedMoney := dao.GetOpenFieldNeedMoney(newFieldId)
		openFieldNeedLevel := dao.GetOpenFieldNeedLevel(newFieldId)
		if userLevel >= openFieldNeedLevel && userMoney >= openFieldNeedMoney {
			dao.OpenField(userName, uId, newFieldId)
			return true
		} else {
			return false
		}

	} else {
		return false
	}
}
