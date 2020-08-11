package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello, ", msg)
}

func (p Person) PrintInfo() {
	fmt.Println("name:", p.Name, "age:", p.Age, "sex:", p.Sex)
}

func (p Person) Test(i, j int, s string) {
	fmt.Println(i, j, s)

}

func main() {
	// reflect对方法的调用
	// step1: 接口变量  ==> 对象反射对象
	// step2 获取对应的方法对象: MethodByName
	// step3: 将方法对象进行调用  Call()

	p1 := Person{"chiling", 48, "男"}
	value := reflect.ValueOf(p1)
	fmt.Println("kind:", value.Kind(), "type: ", value.Type()) // kind: struct type:  main.Person

	// 获取方法
	methodVlue1 := value.MethodByName("PrintInfo")
	fmt.Println("kind: ", methodVlue1.Kind(), "type: ", methodVlue1.Type()) // kind:  func type:  func()

	// 没有参数,进行调用
	methodVlue1.Call(nil) // p1.PrintInfo()  这个里面没有传参的
	//或者传一个空的切片也可以
	args := make([]reflect.Value, 0)
	methodVlue1.Call(args)

	methodVlue2 := value.MethodByName("Say")
	fmt.Println("kind: ", methodVlue2.Kind(), "type: ", methodVlue2.Type()) // kind:  func type:  func(string)  这个有参数了,string
	args2 := []reflect.Value{reflect.ValueOf("hello, world")}
	methodVlue2.Call(args2)

	methodVlue3 := value.MethodByName("Test")
	fmt.Println("kind: ", methodVlue3.Kind(), "type: ", methodVlue3.Type()) // kind:  func type:  func(int, int, string)
	args3 := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf("xxoo")}
	methodVlue3.Call(args3)

}
