package main

import (
	"fmt"
	"math/rand"
	"time"
)
var s = make([]int,0)
func Order (s []int)[]int{
	for i:= 0;i<len(s);i++{
		for j:=i+1;j<len(s);j++{
			if s[i]<s[j]{
				temp := s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}
	}
	return s
}
func main() {
rand.Seed(time.Now().Unix())
for i:=0;i<100;i++{
	var a = rand.Intn(1000)
	s = append(s,a)
}
	fmt.Println(s)
    Order(s)
    fmt.Println("降序后：",s)
}
