package dao

import (
	"farm/app/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

func RegisterFarm(uId string, farmFieldInfoStruct model.FarmFieldInfoStruct) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", append(g.Slice{uId + "_field_" + gconv.String(farmFieldInfoStruct.FieldId)}, gutil.StructToSlice(farmFieldInfoStruct)...)...)

	if err != nil {
		panic(err)
	}
}
func OpenField(userName, uId string, fieldId int) {
	var (
		ctx = gctx.New()
	)

	keyString := uId + "_field_" + gconv.String(fieldId)
	_, err := g.Redis("data").Do(ctx, "HMSET", keyString, "Status", 1)
	if err != nil {
		panic(err)
	}
	_, err = g.Redis("data").Do(ctx, "HMSET", userName, "FieldCount", fieldId)
	if err != nil {
		panic(err)
	}
}

func FieldInfoById(uId string, fieldId int) model.FarmFieldInfoStructWithoutUserId {
	var (
		ctx = gctx.New()
	)

	keyName := uId + "_field_" + gconv.String(fieldId)
	result, err := g.Redis("data").Do(ctx, "HMGET", keyName, "FieldId", "Status", "PlantId", "MaturationTime", "ReMature")
	if err != nil {
		panic(err)
	}
	copyKeyName := uId + "_fieldcopy_" + gconv.String(fieldId)
	fileCopyExist, err := g.Redis("data").Do(ctx, "exists", copyKeyName)
	if err != nil {
		panic(err)
	}
	resultArray := result.Array()
	plantInfo := GetPlantInfo(gconv.Int(resultArray[2]))
	fieldInfo := model.FarmFieldInfoStructWithoutUserId{
		FieldId:        gconv.Int(resultArray[0]),
		Status:         gconv.Int(resultArray[1]),
		PlantName:      plantInfo.PlantName,
		MaturationTime: gconv.Int(resultArray[3]),
		ReMature:       gconv.Int(resultArray[4]),
	}

	//如果土地副本不存在，并且土地状态是2，说明是之前在生长，现在成熟了。
	if gconv.Int(fileCopyExist) == 0 && fieldInfo.Status == 2 {
		//此时需要先把状态改为已成熟
		fieldInfo.Status = 3
		//成熟次数+1次
		fieldInfo.ReMature = fieldInfo.ReMature + 1
		//再判断作物可否再次生长,再次成熟时间
		result, err = g.Redis("config").Do(ctx, "HMGET", fieldInfo.PlantId, "ReMatureCount")
		plantReMatureCount := gconv.Int(result.Array()[0]) + 1
		//plantReMature := gconv.Int(result.Array()[1])
		//已成熟次数小于可以再成熟次数
		if fieldInfo.ReMature < plantReMatureCount {
			//说明可以再次成熟，所以status要写成9
			fieldInfo.Status = 9

		}
		//将结果反写回数据库中，防止后续查询再走一次这个逻辑
		_, err = g.Redis("data").Do(ctx, "HMSET", keyName, "MaturationTime", 0, "Status", fieldInfo.Status, "ReMature", fieldInfo.ReMature)
		if err != nil {
			panic(err)
		}
	}
	return fieldInfo

}

func GetFieldInfo(uId string) [18]model.FarmFieldInfoStructWithoutUserId {
	fieldInfoList := [18]model.FarmFieldInfoStructWithoutUserId{}
	for i := 1; i <= 18; i++ {
		fieldInfo := FieldInfoById(uId, i)
		fieldInfoList[i-1] = fieldInfo
	}
	return fieldInfoList
}

func FieldCopy(uId string, fieldId, plantMaturationTime int) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "SET", uId+"_fieldcopy_"+gconv.String(fieldId), plantMaturationTime)
	if err != nil {
		panic(err)
	}
	_, err = g.Redis("data").Do(ctx, "EXPIRE", uId+"_fieldcopy_"+gconv.String(fieldId), plantMaturationTime)
	if err != nil {
		panic(err)
	}
}
func FieldPlant(uId, userName string, fieldId int, plantInfo model.BagCountStruct, maturationTime int) {
	var (
		ctx = gctx.New()
	)

	_, err := g.Redis("data").Do(ctx, "HMSET", uId+"_field_"+gconv.String(fieldId), "Status", 2, "PlantId", plantInfo.PlantId, "MaturationTime", maturationTime, "ReMature", 0)
	if err != nil {
		panic(err)
	}
}

func FieldHarvest(uId string, fieldId int) {
	//收获后更改土地状态
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", uId+"_field_"+gconv.String(fieldId), "Status", 1, "PlantId", 0, "MaturationTime", 0, "ReMature", 0)
	if err != nil {
		panic(err)
	}
}
func FieldUpgrade(uId string, fieldId, maturationTime, reMature int) {
	//收获后更改土地状态
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", uId+"_field_"+gconv.String(fieldId), "Status", 2, "MaturationTime", maturationTime, "ReMature", reMature)
	if err != nil {
		panic(err)
	}
}
func FieldEradicate(uId string, fieldId int) {
	//收获后更改土地状态
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", uId+"_field_"+gconv.String(fieldId), "Status", 1, "PlantId", 0, "MaturationTime", 0, "ReMature", 0)
	if err != nil {
		panic(err)
	}
	_, err = g.Redis("data").Do(ctx, "DEL", uId+"_fieldcopy_"+gconv.String(fieldId))
	if err != nil {
		panic(err)
	}
}
