/*
@Time : 2022/8/29 09:52
@Author : fushisanlang
@File : textscript
@Software: GoLand
*/
package service

import (
	"farm/app/dao"
)

func GetTextScript(textName string) string {
	TextScript := dao.GetTextScript(textName)
	return TextScript
}
