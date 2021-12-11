package main

import (
	"LV5/Account"
	"LV5/GetPets"
	"LV5/Interface"
	"LV5/Users"
	"fmt"
)

func main() {
	Account.SetFile()  //创建储存玩家账号和密码的txt文本
	Account.SetMassageFile() //创建txt文件SetDownMassage用于写入玩家信息
	for {
		Interface.FirstInterface() //初始登陆界面的显示
		//a refers to the account
		//p refers to the password
		var a, p, c string
		_,_ = fmt.Scanln(&c)
		switch c {
		case "1":
			{
				fmt.Println("Please write you Account here:")
				_,_ = fmt.Scanln(&a)
				fmt.Println("Please write you Password here:")
				_,_ = fmt.Scanln(&p)
				Account.CreatAccount(a, p)   //进行账号和密码的联系
				Account.SetDownAccount(a, p) //把玩家账号和密码储存进UserAccountMassage文件中
			}
		case "2":
			{
				fmt.Println("Please write you Account here:")
				_,_ = fmt.Scanln(&a)
				fmt.Println("Please write you Password here:")
				_,_ = fmt.Scanln(&p)
				temp := Account.Contrast(a, p) //登录时判断账号和密码是否正确并返回调用数据的值
				if temp !=""{
					T := Account.DataLoad(temp)    //加载（读取储存的玩家信息）返回是否成功的标志
					var _U Users.User
					if T && _U.Name==""{
						var U = new(Users.User)
						var GP = new(GetPets.GetPet)
						fmt.Println("选择您的角色(填序号)\n1.小胜 2.小优")
						var m1 int
						_,_ = fmt.Scanln(&m1)
						U.GetName(m1) //选择初始人物
						fmt.Println("选择您的伙伴(填序号)\n1.炎兔儿 2.敲音猴 3.泪眼蜥")
						var m2 int
						_,_ = fmt.Scanln(&m2)
						GP.GetInitialPet(m2) //选择初始伙伴
						Account.SetDownMassage(U,temp) //储存玩家的信息
					}
				}
			}
		case "3":
			{
				fmt.Println("期待下一次见面")
				return
			}
		default:
			fmt.Println("请输入合法选选择")
		}
	}
}