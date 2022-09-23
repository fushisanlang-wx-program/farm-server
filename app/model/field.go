/*
@Time : 2022/8/29 16:24
@Author : fushisanlang
@File : user
@Software: GoLand
*/
package model

type FarmFieldInfoStruct struct {
	UserId         string //用户iD
	FieldId        int    //土地Id
	Status         int    // 是否开启？ 0 关闭 1 开启 2 种植 3 成熟 9 可再次生长
	PlantId        int    //作物id
	MaturationTime int    // 成熟时间
	ReMature       int    //是否再次成熟 0 未成熟过 1 成熟一次 2 成熟二次 3 成熟三次
}

type FarmFieldInfoStructWithoutUserId struct {
	FieldId        int    //土地Id
	Status         int    // 是否开启？ 0 关闭 1 开启 2 种植 3 成熟 9 可再次生长
	PlantName      string //作物id
	MaturationTime int    // 成熟时间
	ReMature       int    //是否再次成熟 0 未成熟过 1 成熟一次 2 成熟二次 3 成熟三次
}
type HarvestInfoStruct struct {
	PlantName  string
	PlantLevel int
	AddEx      int
	AddMoney   int
}
