package service

import (
	"farm/app/dao"
	"farm/app/model"
)

func registerBag(UserName, uId string) {
	for i := 1; i <= 20; i++ {
		BagStruct := model.BagStruct{
			UserId:     uId,
			BagId:      i,
			GoodsId:    0,
			GoodsCount: 0,
		}
		dao.RegisterBag(uId, BagStruct)

	}
}
