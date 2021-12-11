package Users

import (
	"LV5/GetPets"
	"fmt"
)
type User struct {
	string //忽略检索所用的字符串
	Name string
	GetPets.GetPet
}
func (U *User)GetName (a int){
	var F =true
	for F {
		switch a {
		case 1:
			{
				fmt.Println("已选择小胜")
				U.Name = "小胜"
				F = false
			}
		case 2:
			{
				fmt.Println("已选择小优")
				U.Name = "小优"
				F = false
			}
		default:
			{
				fmt.Println("暂无该人物")
			}
		}
	}
}
