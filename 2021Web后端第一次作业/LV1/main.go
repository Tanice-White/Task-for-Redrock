package main

import "fmt"

func main() {
    Str := make(map[int]byte)
    var str []byte
	_,_ = fmt.Scanln(&str)
	for i:=0;i<len(str);i++{
		Str[i] = str[i]
	}
	for _,v := range Str{
		defer fmt.Printf("%v",string(v))
	}

}
