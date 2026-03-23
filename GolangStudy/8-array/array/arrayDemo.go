package main

import "fmt"

func main() {
	var array1 [10]int
	array2 := [10]int{10101, 222, 333}
	for i := 0; i < len(array1); i++ {
		array1[i] = i
	}

	// 遍历数组方式
	for j := 0; j < len(array1); j++ {
		fmt.Println(array1[j])
	}

	for _, varlue := range array1 {
		fmt.Println("value = ", varlue)
	}

	fmt.Println(array2)

}
