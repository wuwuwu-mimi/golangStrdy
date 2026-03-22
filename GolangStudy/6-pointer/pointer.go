package main

import "fmt"

func main() {
	var a1 = 1000
	var a2 = 1999
	swap(&a1, &a2)

	fmt.Println(a1, a2)
}

func swap(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}

func modify(p *int) {
	*p = 20
}
