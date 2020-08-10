package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 3.14
	// "接口类型变量" == "反射类型对象"
	value := reflect.ValueOf(num) // value

	// "反射类型对象" == "接口类型变量"
	converValue := value.Interface().(float64)
	fmt.Println(converValue)  // 3.14

	// 反射类型对象 ==> 接口类型变量,理解为"强制转换"
	// golang对类型要求非常严格,类型一定要完全符合
	// float64 和*float64


	pointer := reflect.ValueOf(&num)
	converPoint := pointer.Interface().(*float64)
	fmt.Println(converPoint)  // 0xc0000120b8

}
