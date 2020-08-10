package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
	Sex string
}

func (p Person) Say(msg string)  {
	fmt.Println("hello, ", msg)
}

func (p Person) PrintInfo()  {
	fmt.Println(p.Name, p.Age, p.Sex)

}
func main() {
	p1 := Person{"chiling", 48, "女"}
	GetMessage(p1)
}
// 获取Input的信息
func GetMessage(input interface{})  {
	getType := reflect.TypeOf(input) // 先获得input的类型
	fmt.Println("Type: ", getType.Name())  // Type:  Person
	fmt.Println("Kind: ", getType.Kind())   // Kind:  struct

	getValue := reflect.ValueOf(input) //获得input的值
	fmt.Println(getValue)  // {chiling 48 女}

	/*
	step1: 先获取Type对象, reflect.Type
	NumField()
	Field(index)
	step2: 通过Filed()获取每一个Filed的字段
	step3: Interface(),提到对应的value

	 */

	for i:=0;i<getType.NumField();i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()// 获取到类型的值
		fmt.Println(field)  // {chiling 48 女}
		fmt.Println(field.Name) // Name  字段名
		fmt.Println(field.Index)  // [0]  字段的index
		fmt.Println(field.Type)  // string 字段的类型
		fmt.Println(value)  // 字段的值
	}

	// 获取结构体中的方法
	for i:=0;i<getType.NumMethod();i++ {
		method := getType.Method(i)
		fmt.Printf("方法名: %s, 方法类型: %v\n", method.Name, method.Type)
		// 方法名: PrintInfo, 方法类型: func(main.Person)
		// 方法名: Say, 方法类型: func(main.Person, string)
	}



}
