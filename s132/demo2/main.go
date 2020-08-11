package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func fun1()  {
	fmt.Println("fun1 fun ....")
}

func fun2(i int, s string)  {
	fmt.Println("fun2 args ...", i, s)
}

func fun3(i int, s string) string {
	fmt.Println("fun3 args return ...", i, s)
	return s + strconv.Itoa(i)
}
func main() {
	// reflect对函数的调用

	// 函数也可以看成是一个接口类型
	// step1: 函数 ==> 反射对象  Value
	// step2: Kind ==> func
	// step3: Call()


	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("kind: %v, type: %v\n", value.Kind(), value.Type())   // kind: func, type: func()


	value2 := reflect.ValueOf(fun2)
	fmt.Printf("kind: %v, type: %v\n", value2.Kind(), value2.Type())   // kind: func, type: func()

	value3 := reflect.ValueOf(fun3)
	fmt.Printf("kind: %v, type: %v\n", value3.Kind(), value3.Type())   // kind: func, type: func(int, string) string


	// 反射来调用函数
	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("hello")})
	retValue := value3.Call([]reflect.Value{reflect.ValueOf(1),reflect.ValueOf("hello")})
	fmt.Println(reflect.TypeOf(retValue), retValue) // []reflect.Value [hello1]
	fmt.Println(retValue[0].Kind(), retValue[0].Type())  // string string
	s := retValue[0].Interface().(string)
	fmt.Println(s)
	fmt.Printf("%T\n", s)













}
