package dao

import (
	"farm/app/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetPlantInfo(plantId int) model.BagCountStruct {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("config").Do(ctx, "HMGET", "plant_"+gconv.String(plantId), "PlantId", "PlantName", "PlantDescription", "Level", "BuyPrice", "MaturationTime", "ReMature", "ReMatureCount", "SellPrice", "Ex")
	if err != nil {
		panic(err)
	}
	resultAttay := result.Array()
	plantInfo := model.BagCountStruct{
		PlantId:          gconv.Int(resultAttay[0]),
		PlantName:        gconv.String(resultAttay[1]),
		PlantDescription: gconv.String(resultAttay[2]),
		Level:            gconv.Int(resultAttay[3]),
		BuyPrice:         gconv.Int(resultAttay[4]),
		MaturationTime:   gconv.Int(resultAttay[5]),
		ReMature:         gconv.Int(resultAttay[6]),
		ReMatureCount:    gconv.Int(resultAttay[7]),
		SellPrice:        gconv.Int(resultAttay[8]),
		Ex:               gconv.Int(resultAttay[9]),
	}
	return plantInfo
}
