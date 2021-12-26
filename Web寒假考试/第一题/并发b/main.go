package main

import "fmt"

func main() {

	var Channel1 = make(chan int,1)

	go func(x int) {
		Channel1<-x
		fmt.Println("下山的路又堵起了")
		<-Channel1
	}(1)

	close(Channel1)
}
