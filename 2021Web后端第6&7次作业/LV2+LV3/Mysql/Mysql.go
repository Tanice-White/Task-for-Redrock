package Mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {                                                  //用户表
	Id int               `json:"id"`//用户的编号
	Name string         `json:"name"` //表示用户名
	Account string      `json:"account"`//表示账号
	Password string     `json:"password"`//表示密码
	Question1 string    `json:"question1"`//表示密保问题(最多20个字/问，最多三个)
	Answer1 string      `json:"answer1"`//保密问题的答案(最多20个字/问,最多三个)
	Question2 string    `json:"question2"`
	Answer2 string      `json:"answer2"`
	Question3 string    `json:"question3"`
	Answer3 string      `json:"answer3"`
}
type Comment struct {                                               //留言表
	Id int                   `json:"id"`//所有评论对应的顺序编号
	FatherId int             `json:"father_id"`//父序号，便于检索
	Layer int                `json:"layer"`//表示评论在第几层
	Status int               `json:"status"`//对应评论的状态(1表示存在,0表示已被删除)    //一般不直接对数据库执行删除操作
	Poster string            `json:"poster"`//发送留言的人的用户名
	Receiver string          `json:"receiver"`//接受者的用户名
	Content string           `json:"content"`//每个人所发送的留言
	Child []*Comment         `json:"child"`//每个留言的回复
}
var UForAll []User //动态储存所有信息的map值
var COForAll []Comment //动态储存留言
var Sum = 0 //Sum用于统计人数
var CSum= 0 //记录留言条数
var DB *sql.DB

//InitDB :与数据库建立链接
func InitDB()error{
	sqlStr := "root:@tcp(127.0.0.1:3306)/text"
	var err error
	DB, err = sql.Open("mysql",sqlStr)
	//设置数据库连接池的最大连接数
	DB.SetMaxIdleConns(10)
	//设置数据库链接池的最大空闲连接数
	DB.SetMaxIdleConns(10)
	return err
}

//ReadAll :读取数据库信息
func ReadAll()error{
	sqlStr := "select id,name,account,password,question1,answer1,question2,answer2,question3,answer3 from `user`;"
	rows, err := DB.Query(sqlStr)
	if err == nil {
		for rows.Next() {
			var v User
			err = rows.Scan(&v.Id,&v.Name,&v.Account,&v.Password,&v.Question1,&v.Answer1,&v.Question2,&v.Answer2,&v.Question3,&v.Answer3)
			UForAll = append(UForAll,v)  //U的初始化
			if err!=nil{
				return err
			}
			Sum++ //此时Sum就是总人数
		}
	}
	_ =rows.Close()      //记得关闭链接!

	//comment的读取
	sqlStr = "select id,father_id,layer,status,poster,`receiver`,content from `comment`;"
	rows, err = DB.Query(sqlStr)
	if err == nil {
		for rows.Next() {
			var c Comment
			err = rows.Scan(&c.Id,&c.FatherId,&c.Layer,&c.Status,&c.Poster,&c.Receiver,&c.Content)
			if c.Status==1{
				COForAll = append(COForAll,c)  //CO的初始化
				if err!=nil{
					return err
				}
				CSum++ //此时Sum就是总人数
			}
		}
	}
	_ = rows.Close()
	return err
}

//WriteIn :将账号密码和用户名储存进数据库
func WriteIn(u User)error{
	sqlStr := "insert into user(name,account,password,question1,answer1,question2,answer2,question3,answer3) values (?,?,?,?,?,?,?,?,?);"
	_, _err := DB.Exec(sqlStr, u.Name, u.Account, u.Password, u.Question1, u.Answer1, u.Question2, u.Answer2, u.Question3, u.Answer3)
	return _err
}
