package main

import (
	"LV2+LV3/Interface"
	"LV2+LV3/Mysql"
	"fmt"
)

func main() {
	err1 := Mysql.InitDB()              //链接数据库
	err2 := Mysql.ReadAll()             //加载所有的数据
	Interface.StartEngines()            //把所有接口全部挂载
	fmt.Printf("数据库链接:%v\n数据加载:%v",err1,err2)
}

//数据库text中创建相关表格的代码
/*  用户表(再所建text数据库中)
 CREATE TABLE `user` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`name` VARCHAR(60) DEFAULT '',
`account`VARCHAR(60)NOT NULL,
`password` VARCHAR(60) NOT NULL,
`question1` varchar(60) default '',
`answer1` varchar(60) default '',
`question2` varchar(60) default '',
`answer2` varchar(60) default '',
`question3` varchar(60) default '',
`answer3` varchar(60) default '',
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/
/* 留言表(在所建的text数据库中)
CREATE TABLE `comment` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`father_id`BIGINT(20) DEFAULT 0,
`layer` BIGINT(10),
`status` TINYINT(1) DEFAULT 1,
`poster` VARCHAR(20) NOT NULL,
`receiver` VARCHAR(20) NOT NULL,
`content` VARCHAR(1000) DEFAULT '',
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
 */