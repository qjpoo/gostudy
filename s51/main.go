package main

import "fmt"

func main() {

	//slice 它是不存储数据的
	/*在已有的数组上直接创建切片*/

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	s1 := a[:5]
	s2 := a[3:8]
	s3 := a[5:]
	s4 := a[:4]
	s5 := a[3:8]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Printf("%p, %p, %p\n", &a, s1, s2) // 0xc000012190, 0xc000012190 地址是一样的

	// 长度和容量
	fmt.Printf("a len： %d, cap: %d, %v\n", len(a), cap(a), a)     // a len： 10, cap: 10, [1 2 3 4 5 6 7 8 9 0]
	fmt.Printf("s1 len： %d, cap: %d, %v\n", len(s1), cap(s1), s1) // s1 len： 5, cap: 10, [1 2 3 4 5] //容量是头到尾 1-10
	fmt.Printf("s2 len： %d, cap: %d, %v\n", len(s2), cap(s2), s2) // s2 len： 5, cap: 7, [4 5 6 7 8]
	fmt.Printf("s3 len： %d, cap: %d, %v\n", len(s3), cap(s3), s3) // s3 len： 5, cap: 5, [6 7 8 9 0]
	fmt.Printf("s4 len： %d, cap: %d, %v\n", len(s4), cap(s4), s4) // s4 len： 4, cap: 10, [1 2 3 4]
	fmt.Printf("s5 len： %d, cap: %d, %v\n", len(s5), cap(s5), s5) // s5 len： 5, cap: 7, [4 5 6 7 8]

	/*	fmt.Println("--------------------改变数组的元素值-------------------------------------------")
		// array和slice都变化了
		a[4] = 100
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
		fmt.Println(s4)
		fmt.Println(s5)
		fmt.Printf("%p, %p, %p, %p, %p\n", &a, s1, s2, s3, s4) // 0xc000012190, 0xc000012190 地址是一样的
	*/
	/*	fmt.Println("--------------------改变切片的元素值-------------------------------------------")
		// array和slice都变化了
		s1[4] = 100
		fmt.Println(a)
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
		fmt.Println(s4)
		fmt.Println(s5)
		fmt.Printf("%p, %p, %p, %p, %p\n", &a, s1, s2, s3, s4) // 0xc000012190, 0xc000012190 地址是一样的
	*/

	fmt.Println("--------------------切片添加元素-------------------------------------------")
	s1 = append(s1, 1, 1, 1, 1, 1)
	fmt.Println(a) // [1 2 3 4 5 1 1 1 1 1]
	fmt.Println(s1) //[1 2 3 4 5 1 1 1 1 1]
	fmt.Println(s2) //[4 5 1 1 1]
	fmt.Println(s3) // [1 1 1 1 1]
	fmt.Println(s4)
	fmt.Println(s5)

	fmt.Println("--------------------切片添加元素-------------------------------------------")
	s1 = append(s1,2, 2, 2, 2)  // s1扩容了, 容量变大了，新申请了一块内存空间地址
	fmt.Println(a) // [1 2 3 4 5 1 1 1 1 1]
	fmt.Println(s1) // [1 2 3 4 5 1 1 1 1 1 2 2 2 2]
	fmt.Println(s2) // [4 5 1 1 1]
	fmt.Println(s3) // [1 1 1 1 1]
	fmt.Println(s4)
	fmt.Println(s5)
}
