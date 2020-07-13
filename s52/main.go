package main

import "fmt"

func main() {
	/*
	切片是引入类型
	*/

	// 数组
	a1 := [4]int{1,2,3,4}
	a2 := a1
	a2[0] = 100
	fmt.Println(a1, a2)

	// slice
	s1 := []int{1,2,3,4}
	s2 := s1
	s2[0] = 100
	fmt.Println(s1, s2)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)



}
