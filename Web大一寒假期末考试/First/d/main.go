package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func Select(x int)bool{           //进行素数的挑选
	for i:=2;i<x;i++{
		if x%i==0{
			return false            //表示不是素数
		}
	}
	return true
}
func main() {
	for j:=2;j<1000000;j++{
		wg.Add(1)
		go func(x int){
			if Select(x) {         //成立则输出
				fmt.Printf("%d，",x)
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
}