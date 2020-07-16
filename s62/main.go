package main

import "fmt"

// 函数的定义写在main之前，或者之后都可以
func sum(x int) int {
	s := 0
	for i:=1;i<=x;i++{
		s += i
	}
	return s
}
func main() { //程序的入口，是一个特殊的函数
	// function
	/*
	可以避免重复的代码
	增强程序的扩展性
	step1: 声明
	step2: 调用

	func funcName(para type1, para type2) (output1 type1, output2 type2) {
	 return value1, value2 //返回多个值
	}
	A: func 关键字
	B: funcName 函数名
	C: 函数的标志
	D: 参数列表 形式参数用于接收外部传入的参数
	E: 返回值 列表，函数执行后返回给调用处的结果

	函数的调用处

	注意事项：
	A: 函数要先定义，再调用
	B: 函数名不能冲突
	C: main(),是一个特殊的函数，作为程序的入口，由系统自动来调用

	*/
	fmt.Println(sum(100)) // 函数的调用处，注意一定要写小括号，不然就是一个变量
	fmt.Println(sum(10))

}
