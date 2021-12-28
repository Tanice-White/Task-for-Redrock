package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()
	mu.Unlock()
}
//可能报错原因：
//代码首先实现了语句的输出才锁上，仍会出现资源的抢夺，导致锁和开锁的顺序错乱而报错。