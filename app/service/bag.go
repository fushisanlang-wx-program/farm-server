package service

import (
	"farm/app/dao"
	"farm/app/model"
)

func registerBag(uId string) {
	dao.SetBag(uId, 1, 10)
	for i := 2; i <= 50; i++ {
		dao.SetBag(uId, i, 0)
	}
}

func GetBagInfo(uId string) []model.BagCountStruct {
	GetBagInfo := dao.GetBagInfo(uId)
	return GetBagInfo

}
func SetPlantCount(uId string, plantId, countNum int) {
	dao.SetBag(uId, plantId, countNum)
}
func getPlantCount(uId string, plantId int) int {
	PlantCount := dao.GetPlantCount(uId, plantId)
	return PlantCount
}
