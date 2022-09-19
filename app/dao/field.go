package dao

import (
	"farm/app/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

func RegisterFarm(Uid string, FarmFieldInfoStruct model.FarmFieldInfoStruct) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", append(g.Slice{Uid + "_field_" + gconv.String(FarmFieldInfoStruct.FieldId)}, gutil.StructToSlice(FarmFieldInfoStruct)...)...)

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
