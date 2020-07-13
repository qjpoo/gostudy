package main

import "fmt"

func main() {
	/*
			数组：定义
			切片： 变长，引用类型的容器
		len 长度
		cap 容量
		定义切片 var s []int
	*/

	//array
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("%T, %v", arr, arr)

	//slice
	s1 := []int{11, 22, 33, 44}
	s2 := []int{111, 222, 333, 444}
	s1 = append(s1, s2...) // 把一个切片放到另一个切片里面，如果不写这个...，表示把s2这个变量append到s1里面
	fmt.Println(s1)

	//切片
	var sa []int
	fmt.Printf("%T, %v\n", sa, sa)

	sb := []int{1, 2, 3}
	fmt.Printf("%T, %v\n", sb, sb)

	var sc []int
	sc = make([]int, 5, 10)
	fmt.Printf("%T, %d, %d\n", sc, len(sc), cap(sc))

	sd := make([]string, 6, 10) //长度为6，现在存储了4个，最大可以存储为10个，可以超过10个
	sd[0] = "A"
	sd[1] = "B"
	sd[2] = "C"
	sd[3] = "D"
	fmt.Printf("%T, %d, %d\n", sd, len(sd), cap(sd))

	//遍历
	for _, v := range sd {
		fmt.Println(v)
	}

	for i := 0; i < len(sd); i++ {
		fmt.Println(sd[i])
	}

	//使用make函数来创建 slice map chan
	// make(type, len, cap)
	//se := make([]int, 3, 6)

	su := make([]int, 5, 10)
	su = append(su, 1, 2, 3)
	fmt.Println(su) // [0 0 0 0 0 1 2 3]
	su = append(su, 10, 20, 30)
	fmt.Println(su) //[0 0 0 0 0 1 2 3 10 20 30]

	var su1 = []int{}
	su1 = append(su1, 1, 2, 3)
	fmt.Println(su1) // [1 2 3]

	su2 := make([]int, 0, 10)
	su2 = append(su2, 1, 2, 3)
	fmt.Println(su2) // [1 2 3]

}
