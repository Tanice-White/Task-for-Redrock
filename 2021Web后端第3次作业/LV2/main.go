package main

import (
	"fmt"
	"os"
)

func main() {
	file,err1 := os.Create("./plant.txt")
	defer file.Close()
	if err1 != nil{
		fmt.Println("文件创建错误",err1)
		return
	}

	//文件的写入
	var Slice ="I’m not afraid of difficulties and insist on learning programming"
	ByteSlice := []byte(Slice)
	n,err2 := file.Write(ByteSlice)
	if err2 != nil{
		fmt.Println("文件写入错误",err2)
		return
	}
	fmt.Printf("共写入了%d个字符\n",n)

	//文件的读取&打印
	File,_:= os.Open("plant.txt")
	defer File.Close()
	var _byteSlice = make([]byte,n)
	_,_= File.Read(_byteSlice)
	fmt.Println(string(_byteSlice))
}