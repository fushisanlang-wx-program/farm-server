package dao

import (
	"farm/app/model"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func SetBag(uId string, bagId, count int) {
	fmt.Println("---------")
	fmt.Println(uId)
	fmt.Println(bagId)
	fmt.Println(count)
	fmt.Println("---------")

	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", uId+"_bag", bagId, count)
	if err != nil {
		panic(err)
	}
}

func GetBagInfo(uId string) []model.BagCountStruct {
	var (
		ctx = gctx.New()
	)
	keyName := uId + "_bag"
	bagList := make(map[int]int)
	for i := 1; i <= 50; i++ {
		result, err := g.Redis("data").Do(ctx, "HMGET", keyName, i)
		if err != nil {
			panic(err)
		}
		plantCount := gconv.Int(result.Array()[0])
		if plantCount > 0 {
			bagList[i] = plantCount

		}

	}
	lenbagList := gconv.Int(len(bagList))

	bagInfoList := make([]model.BagCountStruct, lenbagList, lenbagList)
	j := 0
	for bagId := range bagList {
		result, err := g.Redis("config").Do(ctx, "HMGET", "plant_"+gconv.String(bagId), "PlantId", "PlantName", "PlantDescription", "Level", "BuyPrice", "MaturationTime", "ReMature", "ReMatureCount", "SellPrice", "Ex")
		if err != nil {
			panic(err)
		}
		resultAttay := result.Array()
		bagInfo := model.BagCountStruct{
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
			Count:            GetPlantCount(uId, gconv.Int(resultAttay[0])),
		}
		bagInfoList[j] = bagInfo
		j++
	}

	return bagInfoList

}
func GetPlantCount(uId string, plantId int) int {
	var (
		ctx = gctx.New()
	)
	result, err := g.Redis("data").Do(ctx, "HMGET", uId+"_bag", plantId)

	if err != nil {
		panic(err)
	}
	count := gconv.Int(result.Array()[0])
	return count
}
