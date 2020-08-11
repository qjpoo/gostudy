package main

import (
	"fmt"
	"reflect"
)

type myint int

func main() {
	// reflect设置变量的值
	// Kind指slice map pointer, struct, interface, string, array, function
	// type Person struct{}  kind指 struct, type是指Person

    // 反射变量对应的Kind方法的返回值是基类型，并不是静态类型
	var x myint = 100
	v:= reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Println(v, t) // 100 main.myint
	fmt.Println(v.Kind(),v.Type())   // int main.myint


	fmt.Println(v.Type())
	fmt.Println(v.CanSet())
	fmt.Println(v.Kind())

	n := v.Interface()
	fmt.Println(n, reflect.TypeOf(n))
}
