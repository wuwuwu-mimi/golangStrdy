package main

import "fmt"

// 万能类型
func myFunc(arg any) {
	// 区分类型

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type")
		fmt.Println("value = ", value)
	}
}

func main() {
	myFunc("afjaskdjfa")
}
