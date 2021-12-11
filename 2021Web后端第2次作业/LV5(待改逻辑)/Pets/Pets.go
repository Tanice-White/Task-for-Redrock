package Pets

//Pet :宠物的基础属性
type Pet interface {
	PetName()//宠物的名字
	PetSkill()//技能
	PetHp()//血量
	PetBlue()//蓝量
}

//炎兔儿的数据

type YTE struct{
	Name string
	Skill map[string]string
	HP int
	Blue int
}
func (Y *YTE)PetName()string{
	Y.Name = "炎兔儿"
	return Y.Name
}
func (Y *YTE)PetSkill()map[string]string{
	Y.Skill = make(map[string]string,3)
	Y.Skill["普通攻击"]="普通攻击"
	Y.Skill["技能1：引以为傲的脚力"]="引以为傲的脚力"
	Y.Skill["技能2：火囊"]="火囊"
	return Y.Skill
}
func (Y *YTE)PetHP()int{
	Y.HP = 100
	return Y.HP
}
func (Y *YTE)PetBlue()int{
	Y.Blue = 125
	return Y.Blue
}

//敲音猴的数据

type QYH struct{
	Name string
	Skill map[string]string
	HP int
	Blue int
}
func (Q *QYH)PetName()string{
	Q.Name = "敲音猴"
	return Q.Name
}
func (Q *QYH)PetSkill()map[string]string{
	Q.Skill = make(map[string]string,3)
	Q.Skill["普通攻击"]="普通攻击"
	Q.Skill["技能1：寄生"]="寄生"
	Q.Skill["技能2：木枝突刺"]="木枝突刺"
	return Q.Skill
}
func (Q *QYH)PetHP()int{
	Q.HP = 95
	return Q.HP
}
func (Q *QYH)PetBlue()int{
	Q.Blue = 150
	return Q.Blue
}

//泪眼蜥的数据

type LYX struct{
	Name string
	Skill map[string]string
	HP int
	Blue int
}
func (X *LYX)PetName()string{
	X.Name = "泪眼蜥"
	return X.Name
}
func (X *LYX)PetSkill()map[string]string{
	X.Skill = make(map[string]string,3)
	X.Skill["普通攻击"]="普通攻击"
	X.Skill["技能1：影身"]="影身"
	X.Skill["技能2：哭哭"]="哭哭"
	return X.Skill
}
func (X *LYX)PetHP()int{
	X.HP = 90
	return X.HP
}
func (X *LYX)PetBlue()int{
	X.Blue = 175
	return X.Blue
}