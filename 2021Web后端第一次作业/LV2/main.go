package main

import "fmt"

var (
	a float32
	ch string
	b float32
)
func Add(a float32,b float32){
	fmt.Println("=",a + b)
}
func Minus(a float32,b float32){
	fmt.Println("=",a - b)
}
func Multiply(a float32,b float32){
	fmt.Println("=",a * b)
}

func main() {
	for  {
		fmt.Scan(&a)
		fmt.Scan(&ch)
		fmt.Scan(&b)
		fmt.Println(a, ch, b)
		switch ch {
		case "+":
			{
			Add(a,b)
			}
		case "-":
			{
			Minus(a,b)
			}
		case "*":
			{
			Multiply(a,b)
			}
		default:
			{
				fmt.Println("输入有误")
			}
		}
	}
}