package main

import "fmt"

func main() {
	const (
		BEIJING  = iota * 10 //0
		SHANGHAI             //1
		SHENZHEN             //2
	)

	fmt.Println("BEJIING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
}
