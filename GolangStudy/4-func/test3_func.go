package main

import "fmt"

func func1(a int, b string) int {
	fmt.Println("a = ", a, "b = ", b)
	return 10
}

func main() {
	response := func1(10, "abc")
	fmt.Println(response)
}
