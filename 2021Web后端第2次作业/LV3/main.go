package main

import "fmt"

func Receiver (x interface{}){
	//这里不用if判断
	switch x.(type){      //x.(type)仅仅用于switch！！！
	case int:
		fmt.Println("他是int类型")
	case string:
		fmt.Println("他是string类型")
	case bool:
		fmt.Println("他是bool类型")
	default:
		fmt.Println("不在判定范围内")
	}
}
func main() {
	Receiver(32)
}

