package service

import (
	"farm/app/dao"
	"farm/app/model"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

func registerFarm(uId string) {
	for i := 1; i <= 18; i++ {
		FarmFieldInfoStruct := model.FarmFieldInfoStruct{
			UserId:         uId,
			FieldId:        i,
			Status:         0,
			PlantId:        0,
			MaturationTime: 0,
			ReMature:       0,
		}
		dao.RegisterFarm(uId, FarmFieldInfoStruct)

	}
}

func FieldInfoById(uId string, fieldId int) model.FarmFieldInfoStructWithoutUserId {
	fieldInfo := dao.FieldInfoById(uId, fieldId)
	return fieldInfo

}
func FieldPlant(uId, userName string, fieldId, plantId int) bool {
	//拿到作物信息
	plantInfo := dao.GetPlantInfo(plantId)
	//判断土地是否可用
	fieldStatus := dao.FieldInfoById(uId, fieldId).Status

	//判断种子是否足够
	count := dao.GetPlantCount(uId, plantId)
	//判断用户登记是否足够
	userEx := getEx(userName)
	userLevel := GetLevel(userEx)

	if count < 1 || fieldStatus != 1 || userLevel < plantInfo.Level {
		return false
	} else {
		//dao.FieldPlant(uId, userName, fieldId, plantId)
		//减少一粒种子
		SetPlantCount(uId, plantId, count-1)
		//现在时间
		now := gconv.Int(time.Now().Unix())
		//成熟所需时间
		plantMaturationTime := plantInfo.MaturationTime * 60
		//成熟时间
		maturationTime := plantMaturationTime + now
		//创建一个field副本，用于确认是否成熟
		dao.FieldCopy(uId, fieldId, plantMaturationTime)
		//更改土地信息
		dao.FieldPlant(uId, userName, fieldId, plantInfo, maturationTime)
		//增加经验
		userEx := getEx(userName)
		ChangeEx(userName, userEx+plantInfo.Ex)
		return true
	}
}
func SetFieldCopy(uId string, fieldId, plantMaturationTime int) {
	dao.FieldCopy(uId, fieldId, plantMaturationTime)
}

func GetFieldInfo(uId string) [18]model.FarmFieldInfoStructWithoutUserId {
	fieldInfo := dao.GetFieldInfo(uId)
	return fieldInfo

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
			dao.ChangeMoney(userName, userMoney-openFieldNeedMoney)
			return true
		} else {
			return false
		}

	} else {
		return false
	}
}

func FieldEradicate(uId, userName string, fieldId int) bool {
	//判断土地可否铲除
	fieldInfo := dao.FieldInfoById(uId, fieldId)
	if fieldInfo.Status == 2 || fieldInfo.Status == 3 || fieldInfo.Status == 9 {
		//可以升级

		//更改土地状态，删除可能存在的field副本
		dao.FieldEradicate(uId, fieldId)

		return true
	} else {
		//不能升级
		return false
	}
}

func FieldUpgrade(uId, userName string, fieldId int) bool {
	//var UpgradeInfoStruct model.UpgradeInfoStruct
	//判断土地可否升级
	fieldInfo := dao.FieldInfoById(uId, fieldId)
	if fieldInfo.Status == 9 {
		//可以升级
		//拿到作物信息
		plantInfo := dao.GetPlantInfo(fieldInfo.PlantId)
		//再成熟时间
		ReMatureTime := plantInfo.ReMature * 60

		//更改土地状态
		dao.FieldUpgrade(uId, fieldId, ReMatureTime, fieldInfo.ReMature+1)

		//创建一个field副本，用于确认是否成熟
		dao.FieldCopy(uId, fieldId, ReMatureTime)

		return true
	} else {
		//不能升级
		return false
	}
}

func FieldHarvest(uId, userName string, fieldId int) (bool, model.HarvestInfoStruct) {
	var harvestInfoStruct model.HarvestInfoStruct
	//判断土地可否收获
	fieldInfo := dao.FieldInfoById(uId, fieldId)
	if fieldInfo.Status == 3 || fieldInfo.Status == 9 {
		//可以收获
		//拿到作物信息
		plantInfo := dao.GetPlantInfo(fieldInfo.PlantId)
		harvestInfoStruct.PlantName = plantInfo.PlantName
		harvestInfoStruct.PlantLevel = fieldInfo.ReMature
		//if fieldInfo.ReMature == 2 {
		//	harvestInfoStruct.PlantName = "上品" + harvestInfoStruct.PlantName
		//}
		//if fieldInfo.ReMature == 3 {
		//	harvestInfoStruct.PlantName = "精品" + harvestInfoStruct.PlantName
		//}
		//if fieldInfo.ReMature == 4 {
		//	harvestInfoStruct.PlantName = "珍品" + harvestInfoStruct.PlantName
		//}
		//if fieldInfo.ReMature == 5 {
		//	harvestInfoStruct.PlantName = "极品" + harvestInfoStruct.PlantName
		//}
		//if fieldInfo.ReMature == 6 {
		//	harvestInfoStruct.PlantName = "绝品" + harvestInfoStruct.PlantName
		//}

		//计算经验 ,经验*成熟次数
		harvestInfoStruct.AddEx = plantInfo.Ex * fieldInfo.ReMature
		//计算金钱基础价格*成熟次数
		harvestInfoStruct.AddMoney = plantInfo.SellPrice * fieldInfo.ReMature
		//更改土地状态
		dao.FieldHarvest(uId, fieldId)
		//用户增加经验
		userEx := getEx(userName)
		ChangeEx(userName, userEx+harvestInfoStruct.AddEx)
		//用户增加金钱
		userMoney := getMoney(userName)
		ChangeMoney(userName, userMoney+harvestInfoStruct.AddMoney)

		return true, harvestInfoStruct
	} else {
		//不能收获
		return false, harvestInfoStruct
	}
}
