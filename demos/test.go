/*
@Time : 2022/9/1 15:15
@Author : fushisanlang
@File : test
@Software: GoLand
*/
package main

import (
	"farm/app/dao"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

func main() {
	roleName := gconv.String(dao.GetRoleInfoByUid("1hkzbf25320cmkvrzo3hm9s100kulbta", "RoleName").Slice())
	fmt.Println(roleName)
}
