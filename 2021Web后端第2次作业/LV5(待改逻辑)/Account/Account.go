package Account

import (
	"LV5/Users"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
var Count =-1//表示总账号密码数量（输入时多了一个”\n“先减去）
// SetFile :创建一个储存玩家账号和密码的txt文本
func SetFile (){
	TureFile, err := os.OpenFile("UserAccountMassage.txt",os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Account data Setting error")
	}
	_=TureFile.Close()
}

//AllAccount ：用于临时储存创建的账号和密码
var AllAccount = make(map[string]string,1)
// CreatAccount :进行账号和密码的联系
func CreatAccount(account string,password string){
	AllAccount["account"] = "Password"
}

// SetDownAccount :把玩家账号和密码储存进UserAccountMassage文件中
func SetDownAccount(a,p string){
	TureFile, err := os.OpenFile("UserAccountMassage.txt",os.O_APPEND,0766)
	if err != nil {
		fmt.Println("SetDown error")
	}
	_,err = TureFile.Write([]byte(a+" "+p+"\n"))
	fmt.Println("创建成功！")
	if err != nil {
		fmt.Println("put data error")
	}

	_ =TureFile.Close()
}

// Contrast :登录时判断账号和密码是否正确并返回调用数据的值
func Contrast(account string,password string)string{
	var temp string
	File,_:= ioutil.ReadFile("UserAccountMassage.txt")
	TrueFile :=string(File)
	AccountFile := strings.Split(TrueFile,"\n")
	//每两个为一组，每次检查一组
	for i:=0;;i+=2{
		if AccountFile[i] == ""{
			break
		}else {
			for _,v:= range AccountFile{
				Count++//计数+1，已提前减去最后一个字符”\n“
				if v==account+" "+password{
					fmt.Println("登录成功！")
					temp = "User" + account + password
				}
			}
		}
	}
	if temp ==""{
		fmt.Println("账号或密码错误")
	}
	return temp
}

//SetMassageFile :创建txt文件SetDownMassage用于写入玩家信息
func SetMassageFile(){
	TureFile, err := os.OpenFile("UserMassage.txt",os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("SetDown error")
	}
	_ =TureFile.Close()
}

// SetDownMassage :储存玩家的宠物等信息
func SetDownMassage(U *Users.User,temp string){
	TureFile, err := os.OpenFile("UserMassage.txt",os.O_APPEND, 0766)
	if err != nil {
		fmt.Println("SetDown error")
	}
	for i := 0; i < 1; i++ {
		_,_ = TureFile.WriteString(temp+"  ")
		AllUserMassage,_ := json.Marshal(U)
		_,_ = TureFile.Write(AllUserMassage)
		_,_ = TureFile.WriteString("\n")
	}
	_=TureFile.Close()
}

// DataLoad :加载（读取储存的玩家信息）返回是否成功的标志
func DataLoad(temp string)bool{
	var T = false
	TureFile, err := os.OpenFile("UserMassage.txt",os.O_RDWR, 0766)
	if err != nil {
		fmt.Println("No previous Data")
	}else {
		Massages := make([]byte, 1024)
		_,_ = TureFile.Read(Massages)
		var U Users.User
		_ = json.Unmarshal(Massages,&U)
		_ =TureFile.Close()
		T = true
	}
	return T
}
