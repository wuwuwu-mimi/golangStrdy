package main

import (
	"fmt"
)

func main() {
	var myMap = map[string]int{
		"China": 1,
	}
	//添加
	myMap["China"] = 1
	myMap["USA"] = 2
	myMap["Japan"] = 3

	// 遍历
	for Key, value := range myMap {
		fmt.Println("key = ", Key)
		fmt.Println("value = ", value)
	}

	//修改
	myMap["USA"] = 0

	//删除
	delete(myMap, "Japan")

	fmt.Println("=====================")

	for Key, value := range myMap {
		fmt.Println("key = ", Key)
		fmt.Println("value = ", value)
	}

}
