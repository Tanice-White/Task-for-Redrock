package ControlPanel

import (
	"LV2+LV3/Mysql"
)
var U = new(Mysql.User) //同时临时储存玩家的信息

//FindUser :检索User的动态存储
func FindUser(account string)(u *Mysql.User){
	for k,v:=range Mysql.UForAll{
		if account==v.Account{
			return &Mysql.UForAll[k]
		}
	}
	return &Mysql.User{}
}
//FindComment :检索Comment的动态储存           需要用户名receiver和留言内容content以及id(需要回复则输入)
func FindComment(c Mysql.Comment)(co *Mysql.Comment){ //1找到了会返回那一句留言的地址，0则直接寻找用户名发送留言。否则返回空地址
	if c.Id==0{
		for _,v:=range Mysql.UForAll{
			if c.Receiver==v.Name{   //找到留言对象
				return &Mysql.Comment{
					Id: Mysql.CSum+1,
					Layer: 1,
					Poster: U.Name,
					Receiver: c.Receiver,
					Content: c.Content,
				}
			}
		}
	}else{
		for k,v := range Mysql.COForAll{
			if v.Id==c.Id {
				return &Mysql.COForAll[k]
			}
		}
	}
	return &Mysql.Comment{}      //没有找到留言对象
}

//Contrast :注册时对比账号,密码(不小于6位),用户名是否满足要求    并且对所有账号检索
func Contrast(u Mysql.User)string{
	if Mysql.Sum!=0{
		for _,v:=range Mysql.UForAll{
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

//ToLogIn :登录时检测
func ToLogIn(u Mysql.User)string{
	if Mysql.Sum==0{
		return "没有注册记录"
	}else{
		_u := FindUser(u.Account) //返回的是指向动态总数据的指针
		if _u.Account==u.Account && _u.Password==u.Password{
			U = _u  //同步账号的所有信息
			return "登录成功"
		}
		return "账号或密码错误,是否找回密码？"
	}
}

//ToChangePassword :实现更改密码
func ToChangePassword(u Mysql.User)(err error,str string){
	if len(u.Password)>=6{
		//符合要求则更新数据
		_,err=Mysql.DB.Exec("update user set password=? where account=?;",u.Password,U.Account)
		if err ==nil{
			//同步信息
			U.Password = u.Password
			str = "修改成功，请重新登录"
		}   //有错误则不做操作
	}else{
		str = "密码不符合要求"
	}
	return err,str
}

//ToAnswerQ :比对答案的正确性 并且返回密码
func ToAnswerQ(u Mysql.User)string{
	if U.Question1 ==""{
		return "您未设置密保，无法找回密码"
	}
	//没有题目则答案默认返回空
	if U.Answer1==u.Answer1&&U.Answer2==u.Answer2&&U.Answer3==u.Answer3{
		return "您的密码是:"+U.Password
	}else{
		return "回答错误,无法找回密码"
	}
}

//ToSetQuestions :实现密保的设置
func ToSetQuestions(u Mysql.User)(err error,str string){
	if len(u.Question1)>60||len(u.Question2)>60||len(u.Question3)>60||len(u.Answer1)>60||len(u.Answer2)>60||len(u.Answer3)>60{
		str = "设置不符合要求,请返回检查"
	}else{                           //判断的时候信息已经同步完成了
		if u.Question1!=""{
			U.Question1 = u.Question1
			if u.Answer1!=""{
				U.Answer1 = u.Answer1
			}else{
				return err,"答案不能为空"
			}
		}
		if u.Question2!=""{
			U.Question2 = u.Question2
			if u.Answer2!=""{
				U.Answer2 = u.Answer2
			}else{
				return err,"答案不能为空"
			}
		}
		if u.Question3!=""{
			U.Question3 = u.Question3
			if u.Answer3!=""{
				U.Answer3 = u.Answer3
			}else{
				return err,"答案不能为空"
			}
		}
		sqlStr:="update user set question1=?,answer1=?,question2=?,answer2=?,question3=?,answer3=? where account=?;"
		_,err = Mysql.DB.Exec(sqlStr,U.Question1,U.Answer1,U.Question2,U.Answer2,U.Question3,U.Answer3,U.Account)
		if err != nil{
			return err,str
		}//成功则等待最后的返回
		str = "设置成功!"
	}
	return err,str
}

//ToPostComments :实现留言的发送
func ToPostComments(co Mysql.Comment)(err error,str string){
	x:=FindComment(co)                                         //检索对方的留言的id和用户名找到相应的句子
	if co.Id != 0{    //此时找到了回复目标的地址x                    //找到了需要回复的那句话
		if x.Status ==   0{                                      //表示留言已被删除
			return err,"该留言已被删除"
		}else{
			x.Layer+=1
			sqlStr:="insert into comment (father_id,layer,poster,receiver,content) values (?,?,?,?,?);"
			_,err = Mysql.DB.Exec(sqlStr,x.Id,x.Layer,U.Name,x.Poster,co.Content)
			if err==nil{
				var t = Mysql.Comment{
					Id: x.Id,
					Layer: x.Layer,
					Poster: U.Name,
					Receiver: x.Poster,
					Content: co.Content,
				}
				Mysql.COForAll=append(Mysql.COForAll,t)//同步信息
				str = "留言成功!"
				return err,str
			}else{
				return err,str
			}
		}
	}else{
		sqlStr:="insert into comment (layer,poster,receiver,content) values (?,?,?,?);"
		_,err = Mysql.DB.Exec(sqlStr,x.Layer,x.Poster,x.Receiver,x.Content)
		if err==nil{
			Mysql.COForAll = append(Mysql.COForAll,*x)          //同步信息
			Mysql.CSum++
			return err,"留言成功!"
		}
	}
	return err,"未知错误"
}

//ToChangeComments :更改特定的留言 获取order和change
func ToChangeComments(co Mysql.Comment)string{    //根据内容更改留言
	x:=FindComment(co)                            //检索对方的留言的id找到相应的句子，返回ta的留言的地址
	if x.Poster != U.Name {                       //确定是自己的留言
		return "未找到你的的留言"
	}
	for _,v:= range Mysql.COForAll{
		if co.Id == v.Id {                             //表示留言未被删除
			sqlStr:="update comment set content=? where id=?"
			_,err := Mysql.DB.Exec(sqlStr,co.Content,co.Id)
			if err==nil{
				x.Content = co.Content                 //数据同步
				return "更改成功!"
			}else{
				return "写入错误"
			}
		}else{
			return "留言不存在"
		}
	}
	return "未知错误"
}

//ToDeleteComments :删除留言                         //根据order删除留言
func ToDeleteComments(co Mysql.Comment)string{
	x:=FindComment(co)                              //检索自己的留言的id找到相应的句子，返回自己留言的地址
	if x.Poster != U.Name {                         //确定是自己的留言
		return "未找到你的的留言"
	}else{
		for _,v:= range Mysql.COForAll{
			if co.Id == v.Id {                             //表示留言未被删除
				sqlStr:="update comment set status=0 where id=?;"
				_,err:=Mysql.DB.Exec(sqlStr,co.Id)
				x.Content=""                               //同步信息
				if err!=nil{
					return "数据操作错误"
				}
				Mysql.CSum--
				return "删除成功"
			}else{
				return "留言不存在"
			}
		}
	}
	return "未知错误"
}

//ToShowComments :展示留言
func ToShowComments(n int)(co Mysql.Comment){
	return Mysql.COForAll[n]
}


