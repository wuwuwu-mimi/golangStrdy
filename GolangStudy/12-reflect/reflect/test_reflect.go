package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}
type Employee struct {
	id   int
	name string
}

func (u *User) Call() {
	fmt.Println("user is call ")
	fmt.Printf("u: %v\n", u)
}

// 反射基本类型
func reflectNum(arg any) {
	fmt.Printf("reflect.TypeOf(arg): %v\n", reflect.TypeOf(arg))
	fmt.Printf("reflect.TypeOf(arg): %v\n", reflect.TypeOf(arg))
}

// 反射复杂类型 之 获取 导出字段
func getFieldAndMethond(class any) {
	//获取 input 的 type
	classType := reflect.TypeOf(class)
	fmt.Printf("classType: %v\n", classType)
	//获取class 的 value
	classValue := reflect.ValueOf(class)
	fmt.Printf("classValue: %v\n", classValue)

	// 通过类型获取 类型获取 字段数量
	// 通过字段数量遍历
	for i := 0; i < classType.NumField(); i++ {
		fieldType := classType.Field(i)          // 获取字段类型
		value := classValue.Field(i).Interface() // 获取字段值

		fmt.Printf("字段名 ： %s , 字段类型 : %v , 字段值 ： %v\n", fieldType.Name, fieldType.Type, value)
	}
	fmt.Println("========获取方法=========")
	for i := 0; i < classType.NumMethod(); i++ {
		m := classType.Method(i)
		fmt.Printf("方法名 :%s ---- 类型%v\n", m.Name, m.Type)
	}

}

// 反射复杂类型之 获取 未导出字段
func getNoFiled(class any) {
	t := reflect.TypeOf(class)  //获取反射类型
	v := reflect.ValueOf(class) // 获取反射类型的值 整个结构体

	// 遍历所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		switch fieldValue.Kind() {
		case reflect.Int:
			fmt.Printf("字段：%s，值：%d\n", field.Name, fieldValue.Int())
		case reflect.String:
			fmt.Printf("字段：%s，值：%s\n", field.Name, fieldValue.String())
		}
	}
}

func main() {
	user := User{1, "Ace", 100}
	getFieldAndMethond(&user)
	fmt.Println("============================")
	employee := Employee{1, "zhangsan"}
	getNoFiled(employee)
}
