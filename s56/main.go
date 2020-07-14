package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	map的遍历
	range
	array index value
	map  key value
	*/

	map1 := make(map[int]string)
	map1[1] = "qujian"
	map1[2] = "chiling"
	map1[3] = "linux"
	map1[4] = "golang"
	map1[5] = "C++"
	map1[6] = "ruby"

	// 打印出来是无序的
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println("---------------------")
	//如果key有序的话
	for i := 1; i <= len(map1); i++ {
		fmt.Println(map1[i])
	}

	fmt.Println("---------------------")
	// 如果key是没有序的话
	/*
	step1. 获取所有的key,放到slice里面
	step2. 进行排序
	step3. 遍历key
	*/
	// step1
	s := make([]int,0, len(map1))
	for k, _ := range map1 {
		s = append(s, k)
	}
	fmt.Println(s)
	// step2
	sort.Ints(s)
	fmt.Println(s)
	// step3
	for _, v :=range s {
		fmt.Println(map1[v])
	}

	// 字符串也是可以通过 sort.Strings 来排序
	s1 := []string{"Apple","windows","abc", "aac","abce", "好"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)


}
