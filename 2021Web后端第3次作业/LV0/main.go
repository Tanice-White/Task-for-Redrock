package main

import (
	"fmt"
	"sync"
)
var x int64
var wg sync.WaitGroup
var ch chan int64

func add(ch chan int64) {
	for i := 0; i < 50000; i++ {
		x = x + 1
		ch <- x
	}
	close(ch)
	wg.Done()
}

func MyPrint(ch chan int64){
	for v := range ch{
		fmt.Printf("%d",v)
	}
}
func main() {
	ch = make(chan int64,100)
	wg.Add(2)
	go add(ch)
	go MyPrint(ch)
	wg.Wait()
    //这个打印的太长了叭qwq
}
