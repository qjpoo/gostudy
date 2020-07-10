package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//输入和输出
	// 输出： Print（）
	// Printf() 格式化打印
	// Println()打印换行
	a := 100
	fmt.Printf("%v, %#v, %T", a, a, a)

	b := 3.14
	c := true
	d := "hello"
	e := `Ruby`
	f := 'A'
	fmt.Printf("%T\t%f\n", b, b)
	fmt.Printf("%T\t%t\n", c, c)
	fmt.Printf("%T\t%s\n", d, d)
	fmt.Printf("%T\t%s\n", e, e)
	fmt.Printf("%T\t%d\n", f, f)

	fmt.Printf("%T\t%v\n", b, b)
	fmt.Printf("%T\t%v\n", c, c)
	fmt.Printf("%T\t%v\n", d, d)
	fmt.Printf("%T\t%v\n", e, e)
	fmt.Printf("%T\t%v\t%#v\n", f, f, f)



	// 输入
	// Scanf()
/*	var x int
	var y float64
	fmt.Println("请输入一个整数， 一个浮点类型： ")
	fmt.Scanln(&x, &y)
	fmt.Printf("x的数值为： %d , y的数值为： %f", x, y)
*/
/*	var m int
	var n string
	fmt.Scanf("%d,%s", &m, &n) // 10,string 要按这样的格式输入
	fmt.Printf("数值为：%d , 字符串为：%s", m, n)
*/

	// bufio
	fmt.Printf("请输入一个字符串： ")
	reader := bufio.NewReader(os.Stdin)
	s1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入有问题 ...")
	}
	fmt.Println("读到的数据为： ", s1)



}
