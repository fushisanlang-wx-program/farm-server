package model

type PlantInfoStruct struct {
	PlantId          int
	PlantName        string
	PlantDescription string
	Level            int
	BuyPrice         int
	MaturationTime   int // 成熟时间
	ReMature         int //是否再次成熟 0 未成熟过 1 成熟一次 2 成熟二次 3 成熟三次
	ReMatureCount    int //重复成熟次数
	SellPrice        int
	Ex               int
}
