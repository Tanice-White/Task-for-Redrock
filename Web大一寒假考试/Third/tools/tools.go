package tools

import (
	"Third/global"
)

//FindUser :检索User的动态存储
//参数：登录时的账号
//返回值：特定用户的信息地址
func FindUser(account string)(u *global.User){
	for k,v:=range global.AllUsers{
		if account==v.Account{
			return &global.AllUsers[k]
		}
	}
	return nil
}

//CanRegister :注册时对比账号,密码(不小于6位),用户名是否满足要求    并且对所有账号检索
//参数:注册时获取的账号和密码（使用User结构体传输）
//返回值:注册是否成功的string语句
func CanRegister(u global.User)string{
	if global.Sum!=0{
		for _,v:=range global.AllUsers{
			if u.Account==v.Account || len(u.Password)<6 || u.Name==v.Name{
				return "账号或密码或用户名不符合要求"
			}
		}
		return "注册成功"
	}else{  //之前没有注册，则只检查密码
		if len(u.Password)<6{
			return "密码不符合要求"
		}
		return "注册成功"
	}
}

//CanLogIn :登录时检测
//参数:账号、密码
//	操作：将登录成功的用户信息的地址放入global.ThisUser中
//返回值:登录成功与否的string语句
func CanLogIn(u global.User)string{
	if global.Sum==0{
		return "没有注册记录"
	}else{
		_u := FindUser(u.Account)     //返回的是指向动态总数据的指针
		if _u.Account==u.Account && _u.Password==u.Password{
			global.ThisUser = _u      //同步账号的所有信息
			return "登录成功"
		}
		return "账号或密码错误,是否找回密码？"
	}
}
