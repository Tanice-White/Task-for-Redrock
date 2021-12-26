package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main(){
	for i:= 0;i<10;i++{
		wg.Add(1)   //增加等待时间
		go func() {
			fmt.Println("救命！")
			wg.Done()    //实现完后，时间-1
		}()
	}
	wg.Wait()
}
