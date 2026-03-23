package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4} // 定义动态数组，切片silce

	printArray(array)
	fmt.Println("================")

	for _, value := range array {
		fmt.Println("value = ", value)
	}

}

func printArray(arr []int) {
	for _, value := range arr {
		fmt.Println("value = ", value)
	}
	// 切片是引用传递
	arr[0] = 100
}
