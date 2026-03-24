package main

import "fmt"

//对数据类型声明一个别名
type myint int

//定义一个结构体
type Book struct {
	name string
	age  int
	auth string
}

// 值传递，传递一个副本，修改不会影响到原结构体
func changeBook(book1 Book) {
	book1.age = 1000
}

// 传递一个指针，可以对原结构体修改
func changeBook1(book1 *Book) {
	book1.age = 1000
}

func main() {
	// 使用结构体
	var book1 Book
	book1.age = 12
	book1.auth = "zhangsan"
	book1.name = "母猪的产后护理"
	//changeBook(book1)
	changeBook1(&book1)
	fmt.Println(book1)
	book2 := Book{name: "zhangsan", age: 10, auth: "lisi"}
	fmt.Println(book2)
}
