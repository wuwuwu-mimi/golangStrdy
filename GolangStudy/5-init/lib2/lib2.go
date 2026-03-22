package lib2

import "fmt"

func add(a int, b int) int {
	return a + b
}

func Test() {
	fmt.Println("lib2 ............")
}

func init() {
	fmt.Println("lib2 is running ")

}
