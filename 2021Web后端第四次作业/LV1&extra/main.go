package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Steve struct{
	Skill map[int]string //调用史蒂夫的技能 int用于确定顺序
}
type User struct{
	Skill map[int]string //调用你的技能 int用于确定顺序
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
func main() {
	steve := &Steve{
		Skill: map[int]string{
			0:"龙卷风摧毁停车场",
			1:"乌鸦坐飞机",
			2:"飞鹤捕虾",
			3:"狒朗朗歌舞步",
			4:"蝎子掌",
		},
	}
	var user User
	user.Skill=make(map[int]string,5)
	fmt.Println("请输入你的技能名称(最多储存5个哦)")
	for i:=0;i<5;i++{
		//生成你的技能
		fmt.Printf("请输入第%d个技能的名字,输入end结束录入\n",i+1)
		var str string
		_,_ = fmt.Scanln(&str)
		if str=="end"{
			break
		}
		//开始比对敏感词
		if Contrast(str){
			user.Skill[i]=str
		}else{
			fmt.Println("起名字别这么sao看看？\n回去重启！")
			i = i-1
		}
	}
	//生成一个语句对照表
	var sentence = map[int]string{
		0:"给爷爪巴！！！！",
		1:"你不可能活下来的",
		2:"啊这。。。",
		3:"嘤嘤嘤~~~",
		4:"接招！",
		5:"哇呀呀呀呀呀",
		6:"拿来吧你!",
	}

	fmt.Println("录入完成，下面请开始你的表演吧！")
	fmt.Println("黑老大向你发起了挑战")

	//i用于计数
	i := 0
	for ;i<5;i++{
		var n int
		for k,v := range user.Skill{
			fmt.Printf("技能%d  %s  ",k+1,v)
		}
		fmt.Println("\n发动你的技能！")
		_,_ = fmt.Scanln(&n)
		var f = true
		for k,v := range user.Skill{
			if n-1==k{
				//释放技能并且说一句话
				t := rand.Intn(6)
				str := sentence[t]
				ReleaseSkill(v,func(string){
					fmt.Println("你：",str,"  ",v)
				})
				f = false
				break
			}
		}
		if f{
			fmt.Printf("想啥呢，你哪有这么多技能？？等着被打吧！\n")
		}
		time.Sleep(time.Second)
		//Steve的模块
		for _,v := range steve.Skill{
			//释放技能并且说一句话
			t := rand.Intn(6)
			str := sentence[t]
			ReleaseSkill(v,func(string){
				fmt.Println("黑老大：",str,"  ",v)
			})
		break
	    }

		if i==4{
		fmt.Println("\n黑老大：”啊！我怎么可能输！你个马叉虫！“")
		return
		}
		fmt.Println()
		fmt.Println()
	}
}

//Contrast :敏感词比对函数
func Contrast(str string) bool {
	//context：用于储存文件
	context,err := ioutil.ReadFile("./ku.txt")
	if err != nil {
		fmt.Println("读取失败")
		return false
	}
	Context :=strings.Split(string(context)," ")
	for _,v := range Context {
		fmt.Println(v)
		if strings.Contains(str,v){
			return false
		}
	}
	return true
	//false表示不能添加
}