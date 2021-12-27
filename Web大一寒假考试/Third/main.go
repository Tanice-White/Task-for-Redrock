package main

import (
	"Third/Interface"
	"Third/dao"
	"fmt"
)

func main(){
	err := dao.InitDB()
	if err==nil{
		err = dao.ReadAll()
	}
	Interface.StartEngines()
	fmt.Println(err)
}
//新的数据库(text)中创建相关表格的代码
/*  用户表(再所建数据库中)
 CREATE TABLE `user` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`name` VARCHAR(60) DEFAULT '',
`account`VARCHAR(60)NOT NULL,
`password` VARCHAR(60) NOT NULL,
`money` BIGINT(20) DEFAULT 0,
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/
/* 留言表(在所建的数据库中)
CREATE TABLE `record` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`giver` VARCHAR(60) NOT NULL,
`receiver` VARCHAR(60) NOT NULL,
`money_changed` DECIMAL(10,2),
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/
