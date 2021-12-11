package main

import (
	"fmt"
	"time"
)

var Str =make(map[int]string,11)//只能申请10个闹钟 第11个用来覆盖

var x int //x用于计数

//SetClock :设置闹钟
func SetClock(str string){
	for k,_:= range Str{
		if k+1==11{
			fmt.Println("闹钟数量已达上限")
			return
		}
		if Str[k+1]==""{
			Str[k+1]= str
			x = k+1
		}
	}
	fmt.Println("设置成功！")
}

//DeleteClock :删除闹钟
func DeleteClock(str string){
	for k,v:= range Str{
		if str==v{
			for i:=k;i<10;i++{
				Str[i]=Str[i+1]
				x -=1
			}
		}
	}
	fmt.Println("删除成功！")
}

//CancelClock :取消闹钟的下一次响铃
func CancelClock(str string){
	var p bool
	for _,v := range Str {
		if v==str {
			p = false
			var tmpl = "03:04:05" //输出格式限制
			var t, _ = time.ParseInLocation(tmpl, v, time.Local)
			T := time.Now()
			hour := T.Hour()
			second := T.Second()
			minute := T.Minute()
			if hour == t.Hour() && second == t.Second() && minute == t.Minute() {
				return
			}
		}
	}
	if p{
		fmt.Println("取消失败")
	}else{
		fmt.Println("取消成功")
	}
}

//run :运行闹钟
func run(Str map[int]string){
	for _,v := range Str {
		var tmpl = "03:04:05" //输出格式限制
		var t, _ = time.ParseInLocation(tmpl, v, time.Local)
		T := time.Now()
		hour := T.Hour()
		second := T.Second()
		minute := T.Minute()
		if hour == t.Hour() && second == t.Second() && minute == t.Minute() {
			//协程打开闹钟提醒，并且最多提醒3次
			go func () {
				num := time.NewTimer(time.Minute * 5)
				for range num.C {
					n := 1
					fmt.Println("叮叮叮~闹钟响了")
					n++
					if n == 4 {
						num.Stop()//销毁闹钟的计时
						break
					}
				}
			}()
		}
	}
}

func main() {
	T := time.Now()
	if T.Hour()==02 && T.Minute()==00 && T.Second()==00{
		fmt.Println("谁能比我卷！")
	}
	if T.Hour()==06 && T.Minute()==00 && T.Second()==00{
		fmt.Println("早八算什么，早六才是吾辈应起之时")
	}
	func(){
		t:=time.Tick(time.Second)
		for range t{
			fmt.Println("芜湖！起飞！")
		}
	}()
	func(){
		t:=time.Tick(time.Millisecond*500)
		for range t{
			fmt.Println("芜湖！起飞！")
		}
	}()

	//闹钟的设置
	for {
		if x != 0 {
			run(Str)
			fmt.Println("1.设置闹钟  2.取消闹钟*1  3.删除闹钟  4.保持运行  \n请输入你的选项：")
			var c int
			_,_ = fmt.Scanln(&c)
			switch c {
			case 1:{
				fmt.Println("请输入闹钟的时间")
				var t string
				_,_=fmt.Scanln(&t)
				SetClock(t)
			}
			case 2:{
				fmt.Println("请输入需取消的闹钟的时间")
				var t string
				_,_=fmt.Scanln(&t)
				CancelClock(t)
			}
			case 3:{
				fmt.Println("请输入需删除的闹钟的时间")
				var t string
				_,_=fmt.Scanln(&t)
				DeleteClock(t)
			}
			case 4:{
				fmt.Println("不知道怎么保持程序不做操作了。。。。")
			}
			default:
				fmt.Println("输入错误")
			}
		}
	}
}