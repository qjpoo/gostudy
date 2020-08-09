package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 通道



	var ca chan int
	fmt.Println(reflect.TypeOf(ca), ca)   // chan int <nil>

	ca = make(chan int)
	fmt.Println(reflect.TypeOf(ca), ca)   // chan int 0xc00003e060

	//s := make([]int, 10,10)
	// fmt.Println(reflect.TypeOf(s), s)  //  []int [0 0 0 0 0 0 0 0 0 0]



}
