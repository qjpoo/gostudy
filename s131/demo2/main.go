package main

import (
	"fmt"
	"reflect"
)

func main() {
	// reflect设置值

	var num float64 = 1.23
	fmt.Println("num: ", num)

	// 反射来修改num的值
	// 注意一定是操作的是指针
	pointer := reflect.ValueOf(&num) // 地址
	newValue := pointer.Elem() // 获取指针的value对象x的值
	fmt.Printf("%T\n", newValue)  // reflect.Value
	fmt.Println(newValue)  // 1.23

	fmt.Println("类型： ", newValue.Type())  // 类型：  float64
	fmt.Println("是否可以修改数据： ", newValue.CanSet())   // 是否可以修改数据：  true

	// 重新赋值
	newValue.SetFloat(1.00)
	fmt.Println(newValue)

	// 如果reflect.ValueOf的参数不是指针的话，会panic
	value := reflect.ValueOf(&num)
	value1 := value.Elem()
	value1.SetFloat(1.9) // 如果不是&num,而是num， panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	fmt.Println(value1)
	fmt.Println(value1.CanSet())
	// 如果value不是指针的话 value.Elem()会panic


}
