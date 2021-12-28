package global

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB            //数据库操作
var Engines *gin.Engine   //gin框架操作

var Sum = 0               //用户总人数

type User struct {                       //用户表,包含用户的基本信息
	Id int              `json:"id"`      //用户编号
	Name string         `json:"name"`    //用户名(姓名)
	Account string      `json:"account"` //账号
	Password string     `json:"password"`//密码
	Money   float32     `json:"money"`   //余额
}
var AllUsers []User       //储存所有用户
var ThisUser =new(User)   //储存特定的用户的地址

type RETURN struct {                     //RETURN 接受所有的错误类型
	Err error
	Str string
}
var R RETURN              //储存所有的错误类型

type Record struct {
	Id int              `json:"id"`           //转账的次序
	Giver string        `json:"giver"`        //转账者
	Receiver string     `json:"receiver"`     //接受者
	MoneyChanged float32`json:"money_changed"`//转账金额
}
var AllRecord []Record         //记录所有的交易
var UserRecord Record         //单个用户的交易记录