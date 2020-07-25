package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/*
	错误error
	 */

	f, err := os.Open("./s90")
	fmt.Printf("%T\n", f)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err) // open ./ss90: The system cannot find the file specified.
		return
	}
	fmt.Println(f.Name(), "打开成功 ...")

	/*
	error: 内置的数据类型,内置的接口
	定义方法 : error() string

	可以用errors.New()来创建一个error的对象

	fmt.Errorf()
	*/
	fmt.Println("--------------------------------")
	err1 := errors.New("这是一个错误 ...")
	fmt.Printf("%T\n", err1)
	fmt.Println(err1)
	fmt.Println(err1.Error())

	err2 := fmt.Errorf("错误: %d",100)
	fmt.Println(err2)
	fmt.Printf("%T\n",err2)

	fmt.Println("---------------------")
	erra := checkAge(-3)
	if erra != nil {
		fmt.Println(erra)
		return
	}

}


// 设计一个函数, 验证年龄是否合法,如果为负数,就返回一个error
func checkAge(i int) error{
	if i <= 0 {
		// err := errors.New(" 年纪非法 ...")
		//return  errors.New("age error ...")
		return fmt.Errorf("年纪不合法")
	}
	fmt.Printf("age: %d\n", i)
	return nil
}