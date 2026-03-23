package main

import "fmt"

// defer和return 谁先谁后

func main() {
	ReturnAndDefer()
}
func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func ReturnAndDefer() int {
	defer deferFunc()

	return returnFunc()
}
