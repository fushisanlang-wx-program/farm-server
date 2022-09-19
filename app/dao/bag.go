package dao

import (
	"farm/app/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

func RegisterBag(Uid string, BagStruct model.BagStruct) {
	var (
		ctx = gctx.New()
	)
	_, err := g.Redis("data").Do(ctx, "HMSET", append(g.Slice{Uid + "_bag_" + gconv.String(BagStruct.BagId)}, gutil.StructToSlice(BagStruct)...)...)

	if err != nil {
		panic(err)
	}
}
