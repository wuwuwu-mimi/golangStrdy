package main

// defer的执行顺序

import "fmt"

func main() {

}

func deferA() {
	fmt.Println("A")
}
func deferB() {
	fmt.Println("B")
}
func deferC() {
	fmt.Println("C")
}

// 运行结果为 C B A 先进后出
