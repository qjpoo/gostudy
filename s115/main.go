package main

import (
	"fmt"
)

func main() {
	// goroutines
	// 执行main函数的叫主goroutines,不是main函数的叫子goroutines

	// 1. 先创建并启动一个子goroutines
	// 注意点:
	// 1. 当主的goroutines执行完了,就算子的goroutines还没有执行完成,也不会执行子goroutines了
	// 2. 子goroutines不要有返回值

	// 主goroutes和子的goroutes是交替执行的
	go printNum()

	for i:=1;i<20;i++ {
		fmt.Println("主goroutines: ", i)
	}

	//time.Sleep(time.Second )
	fmt.Println("main over ...")
}
func printNum()  {
	for i:=1;i<50;i++{
		fmt.Println("子goroutines:　", i)
	}

}
