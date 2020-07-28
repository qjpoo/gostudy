package main

import "fmt"

/*
1. defer语句先定义的，最后才会执行
 */

func myPrint(s string)  {
	fmt.Println(s)

}
func funA()  {
	fmt.Println("这是funA函数 ...")
}

// 当外围函数的代码中发生运行了恐慌，只有其中所有的已经defer的函数全部都执行完成了，才会回到函数的调用处
func funB()  { // 外围函数
	defer func() {
		if msg := recover(); msg !=nil {
			fmt.Println(msg, "程序回复啦 。。。")
		}
	}()
	fmt.Println("这是funB函数 ...")
	defer myPrint("defer funB(): 1 ...")  // 先定义，后执行
	for i:=0;i<10;i++ {
		fmt.Println(i)
		if i == 5 {
			// 让程序 panic
			panic("funB() panic ... ") // 后面的代码都不执行了
		}
	}
	defer myPrint("defer funB(): 2 ...")
}
func main() {
	/*
	panic recover
	defer是按定义的逆顺来执行
	 */

	// 等所有的都执行完成了在执行defer
	funA()
	defer myPrint("defer main: 3 ...")
	funB()  // 当funB()发生了panic，后面的 defer和println也不会执行了
	defer myPrint("defer main: 4 ...")

	fmt.Println("main over ...")
}
