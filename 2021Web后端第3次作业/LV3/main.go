package main

import (
	"fmt"
	"sync"
)
var mutex sync.RWMutex
var wg sync.WaitGroup
func main() {
	over := make(chan bool,1) //原码 over := make(chan bool) 没有声明管道的长度
	ch := make(chan int,1) //增加用于读写的管道
	for i := 0; i < 10; i++ {
		/*
		//存在多个协程同时读取i & 协程未执行完的情况的情况
		go func() {
			fmt.Println(i)
		}()
		 */
		ch <- i
		wg.Add(1)
		go func() {
			mutex.Lock()
			fmt.Println(<-ch)
			mutex.Unlock()
			wg.Done()
		}()
		if i == 9 {
			over <- true
		}
	}
	<-over
	//没有关闭通道
	close(over)
	close(ch)
	wg.Wait()
	fmt.Println("over!!!")
}
