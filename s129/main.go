package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 反射 reflect
	// 只有Interface类型才有反射一说

	var x float64 = 3.14

	fmt.Println("type: ", reflect.TypeOf(x))  // type:  float64
	fmt.Println("value: ", reflect.ValueOf(x))  // value:  3.14

	fmt.Println("-----------------------")
	//根据反射的值,来获取对应的类型和数值
	v := reflect.ValueOf(x)
	fmt.Printf("%T, %v\n", v, v)  // reflect.Value, 3.14
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)  // kind is float64:  true   reflect.Float64是常量
	fmt.Println("type: ", v.Type())  // type:  float64
	fmt.Println("value: ", v.Float())  // value:  3.14

	
}
