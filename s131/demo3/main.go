package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {

	s1 := Student{"chiling", 48, "家里墩大学"}
	// 通过反射，更改对象的数值，前提也是数据可以被更改
	fmt.Printf("%T\n", s1) // main.Student
	p1 := &s1
	fmt.Printf("%T, %v\n", p1, p1) // *main.Student
	fmt.Println(s1.Name)
	fmt.Println(p1.Name, (*p1).Name)

	/*
	a := 3
	var pp *int
	pp = &a
	fmt.Printf("%T, %v\n", pp, pp)
	 */

	// 改变数值
	value := reflect.ValueOf(&s1)  // 一定要是地址
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println(newValue.CanSet())   // true

		f1 := newValue.FieldByName("Name")
		f1.SetString("qujian")
		f3 := newValue.FieldByIndex([]int{2}) // 通过index来访问字段School
		f3.SetString("武汉大学")
		fmt.Println(newValue,s1)
	}


}
