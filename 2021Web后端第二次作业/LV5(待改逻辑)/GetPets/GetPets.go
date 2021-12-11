package GetPets

import (
	"LV5/Pets"
	"fmt"
)

type GetPet struct {
	GetPetName string//宠物的名字
	GetPetSkill map[string]string//技能
	GetPetHp int//血量
	GetPetBlue int//蓝量
}
/*
1 : 炎兔儿
2 : 敲音猴
3 : 泪眼蜥
 */

// GetInitialPet :获取初始宠物
func (GP *GetPet)GetInitialPet(a int) {
	var PF = true
	for PF {
		switch a {
		case 1:
			{
				fmt.Println("已选择炎兔儿！")
				var Y =new(Pets.YTE)
				GP.GetPetName = Y.PetName()
				GP.GetPetSkill = Y.PetSkill()
				GP.GetPetHp = Y.PetHP()
				GP.GetPetBlue = Y.PetBlue()
				PF = false
			}
		case 2:
			{
				fmt.Println("已选择敲音猴！")
				var Q =new(Pets.QYH)
				GP.GetPetName = Q.PetName()
				GP.GetPetSkill = Q.PetSkill()
				GP.GetPetHp = Q.PetHP()
				GP.GetPetBlue = Q.PetBlue()
				PF = false
			}
		case 3:
			{
				fmt.Println("已选择泪眼蜥！")
				var X =new(Pets.LYX)
				GP.GetPetName = X.PetName()
				GP.GetPetSkill = X.PetSkill()
				GP.GetPetHp = X.PetHP()
				GP.GetPetBlue = X.PetBlue()
				PF = false
			}
		default:
			{
				fmt.Println("必须选择您的第一只宝可梦哦")
			}
		}
	}
}
