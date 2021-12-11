package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type UserData struct{
	Key string
	Value string
}
//Read :将数据导入user结构体中备用,随时可用于比较
func Read(){
	byteSlice,err:=ioutil.ReadFile("./user.data")
	if err!=nil{
		fmt.Println("文件操作错误-读取")
	}
	//将文件分成不同的组,
	byteString := string(byteSlice)
	StrSlice := strings.Split(byteString,"\n") //会多读出来一个空的Slice
	//每个组放入一个map数据内的结构体中
	var temp UserData
	for k:= range StrSlice {
		//顺便判断一下number的初始值
		if StrSlice[k]!=""{
			number = k+1
		}
	}
	//避免重复赋值
	for i:=0;i<number;i++{
		_ = json.Unmarshal([]byte(StrSlice[i]),&temp)
		user[i]=append(user[i],temp)
	}
}

//WriteIn :将user结构体中的数据导入库中储存，用JSON格式
func WriteIn(data UserData){
	file,err := os.OpenFile("./user.data",os.O_APPEND,0666)
	if err!= nil{
		fmt.Println("文件操作错误-写入",err)
	}
	//将转结构体换为JSON字符串（序列化），再写入文本，
	var used = new(UserData)
	used.Key = data.Key
	used.Value = data.Value //已加密
	ByteStr,_ :=json.Marshal(used)
	_,err = file.Write(ByteStr)
	_,err = file.WriteString("\n")
	if err!=nil{
		fmt.Println("写入错误",err)
	}
	number++
	_ = file.Close()
}

//HMAC 加密函数,用账号加密密码，卷起来
func HMAC(key,data string)string{
	hash := hmac.New(md5.New,[]byte(key))//创建对应md5的加密算法
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

//contrast :登录时验证账号和密码
func contrast(key,data string)bool{
	for _,v:=range user{
		for _,_v := range v{
			if _v.Key == key{
				if _v.Value == HMAC(key,data){
					return true
					//true表示允许登录
				}
			}
		}
	}
	return false
	//false表示账号或密码错误
}
//Require :注册时检测账号和密码是否符合要求
func Require(key,data string)bool{
	if len(data)<6{
		fmt.Printf("密码不符合规则")
		return false
	}
	for _,v:=range user{
		for _,_v := range v{
			if _v.Key == key || _v.Value == HMAC(_v.Key,data){
				fmt.Printf("账号或密码重复")
				return false
			}
		}
	}
	return true   //true表示允许注册
}

//Change :更改密码
func Change(key,value string)bool{
	if len(value)<6{
		fmt.Println("密码不符合要求")
	}else{
		var Temp =UserData{
			key,
			HMAC(key,value),
		}
		WriteIn(Temp) //它会自动加密
		number++
		//把增加的值同步
		user[number]=append(user[number],Temp)
		//查重，删重操作(key相同则删除前面的)
		return accountDelete(key)
		//true 表示更改密码成功
	}
	return false
}
//accountDelete :查重，删重
func accountDelete(account string)bool{
	for k,v := range user {
		for _, _v := range v {
			//找出账号相等的两个账户，消除前一个
			if _v.Key == account {
				for i:=k;i<len(user);i++{
					user[k]=user[k+1]
				}
				number--
			}
		}
	}
	//得到无重复的user数组,将其逐一输入库中
	var arr=make(map[int]string,number+1) //临时储存账号
	var marr=make(map[int]string,number+1) //临时储存密码
	for k,v := range user{
		for _,_v := range v{
			arr[k] = _v.Key
			marr[k] = _v.Value
		}
	}
	_f,err := os.OpenFile("./user.data",os.O_TRUNC|os.O_WRONLY,0666)
	if err!=nil{
		fmt.Println("查重删重时出错了qwq",err)
		return false
	}
	//这个100我找了好久，防止打印不了！！！！！！
	for i:=0;i<100;i++{
		if arr[i]!="" && marr[i]!=""{
			var tStruct =UserData{
				arr[i],
				marr[i],
			}
			WriteIn(tStruct)
		}
	}
	_ =_f.Close()
	return true
}

//全局变量的声明
var number =0 //统计已有的数据信息（账号的个数）
var user =make(map[int][]UserData,number+1) //初始化map的长度（增加的一个用来覆盖）

func main() {
	f,_ :=os.OpenFile("./user.data",os.O_CREATE|os.O_RDWR,0666)
	_ = f.Close()//文件关闭，后面还会打开的
	Read() //预加载储存的内容并赋值给user
	for{
		var a int
		fmt.Println("1.注册 2.登录 3.退出\n请选择你的操作:")
		_,_ = fmt.Scanln(&a)
		switch a {
		//1表示注册账号，要求：用户名不能重复，密码长度大于6位，不允许重复注册
		case 1:{
			fmt.Println("\t—注册入口—\n要求：账号不重复，密码大于6位")
			var account,password string
			fmt.Print("请输入你的账号：")
			_,_ = fmt.Scanln(&account)
			fmt.Print("请输入你的密码：")
			_,_ = fmt.Scanln(&password)
			fmt.Println()
			//验证是否符合规则
			if Require(account,password){
				//实例化结构体临时储存
				var TempUser = UserData{
					account,
					HMAC(account,password), //加密再写入
				}
				WriteIn(TempUser)  //将注册好的信息录入库中储存
				//同步user
				user[number] = append(user[number],TempUser)
				fmt.Println("注册成功，请重新登录")
			}else{
				fmt.Println("\n注册失败")
			}
			break //退出登录，返回初始界面
		}
		//2表示登录账号。
		case 2:{
			if number ==0{
				fmt.Println("您还未创建账户！")
				break //无法登录，返回初始界面
			}
			fmt.Println("\t—登录入口—")
			var account,password string
			fmt.Print("请输入你的账号：")
			_,_ = fmt.Scanln(&account)
			fmt.Print("请输入你的密码：")
			_,_ = fmt.Scanln(&password)
			//判断是否允许登录（contrast会自动加密密码）
			if contrast(account,password){
				fmt.Println("登录成功！")
				for {
					fmt.Println("您已登陆，目前只能进行改密码操作（输入1更改密码，输入end退出系统）：")
					var answer string
					_, _ = fmt.Scanln(&answer)
					switch answer {
					// 更改密码
					case "1":
						{
							fmt.Println("请输入更改后的密码：(大于6位)")
							var tempValue string
							_,_ = fmt.Scanln(&tempValue)
							if Change(account,tempValue){
								fmt.Println("更改成功!")
								break
							}else{
								fmt.Println("更改失败")
								break //回到登陆后的页面
							}
						}
					//退出服务系统
					case "end":
						{
							fmt.Println("感谢使用！")
							return
						}
					default:
						fmt.Println("其他功能暂未开发qwq")
					}
				}
				//不允许时
			}else{
				fmt.Println("账号或密码错误！")
			}
			break //登录失败，返回初始页面
		}
		//退出系统
		case 3:{
			fmt.Println("感谢使用！")
			return
		}
		default:
			fmt.Println("请输入合法的数据")
		}
	}
}