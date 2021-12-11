package main

import (
	"fmt"
	"sync"
)

func Print (){
	for i:=0;i<=100;i++{
		fmt.Printf("%d ",i)
	}
	wg.Done()
}

var wg sync.WaitGroup
func main() {
    wg.Add(1)
    go Print()
    wg.Add(1)
    go Print()
	wg.Wait()
}
