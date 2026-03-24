package main

import (
	"fmt"
)

// interface 本质是一个指针
type AnimalInterface interface {
	Sleep()
	GetClolor() string
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

// 实现接口的方法
func (c *Cat) Sleep() {
	fmt.Println("cat is sleep")
}
func (c *Cat) GetClolor() string {
	return c.color
}
func (c *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

// 实现接口的方法
func (d *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (d *Dog) GetClolor() string {
	return d.color
}
func (d *Dog) GetType() string {
	return "Dog"
}

// 可以传递 dog 也可以是 cat
func showAnimal(animal AnimalInterface) {
	fmt.Printf("animal.GetType(): %v\n", animal.GetType())
}

// 多态的现象
func main() {
	var animal AnimalInterface
	//因为 interface的本质是指针，需要传递给它地址
	animal = &Cat{"Green"}
	// 此时调用的是 cat的sleep 方法
	animal.Sleep()
	fmt.Printf("animal.GetClolor(): %v\n", animal.GetClolor())
	showAnimal(animal)
	animal = &Dog{"Red"}
	// 此时调用的是 dog的sleep方法
	animal.Sleep()
	fmt.Printf("animal.GetClolor(): %v\n", animal.GetClolor())
	showAnimal(animal)
}
