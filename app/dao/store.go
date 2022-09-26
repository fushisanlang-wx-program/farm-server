package dao

import (
	"farm/app/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetStoreList() [50]model.PlantInfoStruct {

	var ctx = gctx.New()

	var storeList [50]model.PlantInfoStruct
	var plantInfo model.PlantInfoStruct
	for i := 1; i < 51; i++ {

		result, err := g.Redis("config").Do(ctx, "HMGET", "plant_"+gconv.String(i), "PlantId", "PlantName", "PlantDescription", "Level", "BuyPrice", "MaturationTime", "ReMature", "ReMatureCount", "SellPrice", "Ex")
		if err != nil {
			panic(err)
		}
		resultArray := result.Array()

		plantInfo.PlantId = gconv.Int(resultArray[0])
		plantInfo.PlantName = gconv.String(resultArray[1])
		plantInfo.PlantDescription = gconv.String(resultArray[2])
		plantInfo.Level = gconv.Int(resultArray[3])
		plantInfo.BuyPrice = gconv.Int(resultArray[4])
		plantInfo.MaturationTime = gconv.Int(resultArray[5])
		plantInfo.ReMature = gconv.Int(resultArray[6])
		plantInfo.ReMatureCount = gconv.Int(resultArray[7])
		plantInfo.SellPrice = gconv.Int(resultArray[8])
		plantInfo.Ex = gconv.Int(resultArray[9])

		storeList[i-1] = plantInfo
	}

	return storeList

}
