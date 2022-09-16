/*
@Time : 2022/8/31 09:47
@Author : fushisanlang
@File : role
@Software: GoLand
*/
package model

type RoleStruct struct {
	RoleName   string
	Jin        int
	Mu         int
	Shui       int
	Huo        int
	Tu         int
	TalentedId int
	CreateDone int
}
type PanelPropertiesStruct struct {
	PropertiesTi   float32 //体：血量
	PropertiesJu   float32 //聚：聚气，选定一个属性进行修炼，缓慢增加属性值。
	PropertiesSu   float32 //速：攻击速度，CD
	PropertiesLi   float32 //力：攻击力
	PropertiesYu   float32 //御：防御力
	PropertiesPo   float32 //破：破甲
	PropertiesMing float32 //命：攻击命中率
	PropertiesShan float32 //闪：闪避机率
	PropertiesFu   float32 //腐：毒
	PropertiesJi   float32 //汲：吸血
	PropertiesFeng float32 //封：控制加强
	PropertiesRen  float32 //韧：控制减弱
}
type TalentedSkillStruct struct {
	TalentedSkillId       int
	TalentedSkillName     string                //:宝藏xx
	TalentedSkillNameInfo string                //：xxxxx
	TalentedSkill         PanelPropertiesStruct //：{力：1，速：-1}
	TalentedSkillInfo     string                //：增加攻击力，增加防御力

}
