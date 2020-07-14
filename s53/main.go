package main

import "fmt"

func main() {
	/*
		深拷贝  值
		浅拷贝  地址
	*/

	// 深拷贝
	var i int
	i = 100
	j := i
	j = 200
	fmt.Println(i, j) // 修改了j的值 ，i的值不会变化

	// 浅拷贝
	var m = []int{1, 2, 3, 4, 5}
	n := m
	n[0] = 100
	fmt.Println(m, n) //修改n的值的时候，m的值也会修改了

	// 使用append方法，给赋值
	n1 := make([]int, 0)
	for i := 0; i < len(m); i++ {
		n1 = append(n1, m[i])
	}
	fmt.Println(n1)

	fmt.Println("----------------------------------")
	a1 := []int{1, 2, 3, 4}
	a2 := []int{11, 22, 33}
	//copy(a2, a1)  // 将源a1放进目标a2里面
	//fmt.Println(a2) // [1 2 3]

	//copy(a1, a2 )
	//fmt.Println(a1, a2) // [11 22 33 4] [11 22 33]

	// copy(a2, a1[1:])
	// fmt.Println(a1, a2) // [1 2 3 4] [2 3 4]

	copy(a2[1:], a1[3:])
	fmt.Println(a1, a2)  //[1 2 3 4] [11 4 33]

}
