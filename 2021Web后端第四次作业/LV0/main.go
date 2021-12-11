package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 目的：查找某文章的单词
func main() {
   file,err := os.OpenFile("./index.txt",os.O_RDWR,0666)
   defer file.Close()
   if err != nil{
	   fmt.Println("打开文件失败！")
	   return
   }
   //context：用于储存文件
   context,_err := ioutil.ReadFile("./index.txt")
	if _err != nil {
		fmt.Println("读取失败")
		return
	}
  //把文件里面的文章读取出来，保证不修改原文
  Context :=string(context)
  var str string
  fmt.Println("请输入需查找的符号：")
  fmt.Scanln(&str)

  //查找和强调的实现
  func(){
	 f := strings.Contains(Context,str)
	 if f{
		 fmt.Printf("本文中%s共出现了%d次\n",str,strings.Count(Context,str))
		 //原本还想输出原文并且强调一下搜索的部分的，可惜不会。。。
	 }else{
		 fmt.Printf("本文并不包含%s",str)
		 return
	 }
  }()
}
