package lib1

import "fmt"

func add(a int, b int) int {
	return a + b
}

func Test() {
	fmt.Println("lib1 ............")
}

func init() {
	fmt.Println("lib1 is running ")
}
