package main

import "fmt"

func sum(x ...int) {
	sum := 0
	fmt.Printf("%T\n", x) // x是一个slice类型的
	for _, v := range x {
		sum += v
	}
	fmt.Printf("%d\n", sum)
}

func	xxoo(a, b int, c ...int) int {
	m := 0
	for i:=0;i<len(c);i++{
		m += c[i]
	}
	return a + b + m
}
func main() {
	// 可变参数
	// 可变参数里面的形参是一个切片类型 Println, Printf, Print都是接收可变参数的， append
	// 注意事项： A: 如果一个函数有可变参数和还有其它的参数的话，可变参数要放在后面。 B：一个函数的可变参数最多只有一个

	sum(10, 20)
	sum(1, 2, 3, 4, 5)

	s := []int{1,2,3,4,5}
	sum(s...)  // ... 获取s切片中的数据

	fmt.Println(xxoo(1,2,3,4,5))

}
