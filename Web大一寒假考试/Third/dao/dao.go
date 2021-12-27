package dao

import (
	"Third/global"
	"Third/tools"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//InitDB :与数据库建立链接
func InitDB()(err error){
	sqlStr := "root:@tcp(127.0.0.1:3306)/text"
	global.DB, err = sql.Open("mysql",sqlStr)
	//设置数据库连接池的最大连接数
	global.DB.SetMaxIdleConns(10)
	//设置数据库链接池的最大空闲连接数
	global.DB.SetMaxIdleConns(10)
	return err
}

//ReadAll :读取数据库信息
//初始化global.Sum总人数和global.AllUser
func ReadAll()error{
	sqlStr := "select id,name,account,password,money from `user`;"
	rows, err := global.DB.Query(sqlStr)
	if err == nil {
		for rows.Next() {
			var tempUser global.User
			err = rows.Scan(&tempUser.Id,&tempUser.Name,&tempUser.Account,&tempUser.Password,&tempUser.Money)
			global.AllUsers = append(global.AllUsers,tempUser)  //用户的初始化
			if err!=nil{
				return err
			}
			global.Sum++                                        //此时Sum就是总人数
		}
	}else{
		_ =rows.Close()
		return err
	}

	sqlStr = "select giver,`receiver`,money_changed from `record`;"
	rows, err = global.DB.Query(sqlStr)
	if err == nil {
		for rows.Next() {
			var tempRecord global.Record
			err = rows.Scan(&tempRecord.Giver,&tempRecord.Receiver,&tempRecord.MoneyChanged)
			global.AllRecord = append(global.AllRecord,tempRecord)  //交易记录的初始化
			if err!=nil{
				return err
			}
		}
	}
	_ =rows.Close()
	return err
}

//WriteIn :将账号密码和用户名储存进数据库
//参数:用户的结构体，包含用户名、账户、密码
//返回值:返回错误
func WriteIn(u global.User)error{
	sqlStr := "insert into user(name,account,password) values (?,?,?);"
	_, _err := global.DB.Exec(sqlStr, u.Name, u.Account, u.Password)
	return _err
}

//HasChangePassword :更改密码
//参数:更改后的密码
//返回值:错误语句err和str
func HasChangePassword(u global.User)(err error,str string){
	if len(u.Password)>=6{
		//符合要求则更新数据
		_,err=global.DB.Exec("update user set password=? where account=?;",u.Password,global.ThisUser.Account)
		if err ==nil{
			//同步信息
			global.ThisUser.Password = u.Password
			str = "修改成功，请重新登录"
		}   //有错误则不做操作
	}else{
		str = "密码不符合要求"
	}
	return err,str
}

//HasCharge :给自己的账户充值
//参数:存入的金额(默认为正)
//返回值:是否充值成功的语句
func HasCharge(moneyGet float32)(err error,str string){
	if moneyGet>0{
		moneyGet += global.ThisUser.Money
		_,err = global.DB.Exec("update user set money=? where account=?;",moneyGet,global.ThisUser.Account)
		if err==nil{
			global.ThisUser.Money = moneyGet
			str = "充值成功！"
		}
	}else{
		str = "金额不合法"
	}
	return err,str
}

//HasTransfer :转账
//参数:转入的账户、转出的金额,打包为用户的结构体
//返回值:是否转账成功的语句
func HasTransfer(u global.User)(err error,str string){
	if tools.FindUser(u.Account)==nil{
		return err,"此用户不存在"
	}
	if u.Account==global.ThisUser.Account{
		return err,"无法给自己转账"
	}
	if global.ThisUser.Money>u.Money{
		global.ThisUser.Money -= u.Money                           //转出者的余额
		_,err = global.DB.Exec("update user set money=? where account=?;",global.ThisUser.Money,global.ThisUser.Account)
		if err == nil{
			tempUser := tools.FindUser(u.Account)                  //找到接收者的信息（地址）
			tempUser.Money += u.Money                              //接收者的余额
			_,err = global.DB.Exec("update user set money=? where account=?;",tempUser.Money,u.Account)
			if err == nil{
				err = HasRecord(u)                                 //生成记录
				if err == nil{
					str = "转账成功！"
				}
			}
		}
	}else{
		str = "余额不足"
	}
	return err,str
}

//HasRecord :生成交易记录
//参数:接收到的用户结构体，包含:转出金额、接收者
//	操作:将交易记录写入数据库
//返回值:错误信息
func HasRecord(u global.User)error{
	sqlStr := "insert into record (giver,receiver,money_changed)values(?,?,?)"
	_,err := global.DB.Exec(sqlStr,global.ThisUser.Account,u.Account,u.Money)
	return err
}
