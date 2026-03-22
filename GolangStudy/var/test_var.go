package main

import "fmt"

func main() {
	var a int = 10

	fmt.Println("a = ", a)
	fmt.Printf("a type of = %T\n", a)

	var str string = "abc"
	fmt.Printf("str = %s  type of str = %T\n", str, str)

	str2 := "abcd"

	fmt.Printf("type of str2 = %T\n", str2)

	var (
		a1 int
		b  int
	)
	fmt.Println("a = ", a1, "b =", b)

}
