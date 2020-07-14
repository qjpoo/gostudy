package main

import "fmt"

func main() {
	/*
	map key value
	zhangs 123
	map是无序的，长度不固定，可以用Len来获取长度
	int 0
	float 0.0
	string: ""
	array: [00]
	map: nil
	slice: nil
	nil 的map是不能使用的
	var map1 map[int]string 只是定义了，但是没有初始化
	通过Make函数来初始化    map1 = make(map[int]string)

	*/

	// map
	var map1 map[int]string // nil
	//map1[1] = "hello" // 只是定义了，但是没有初始化，这样写是不对的
	fmt.Println(map1 == nil)

	var map2 = make(map[int]string)
	var map3 = map[string]int{"go": 91, "java": 90, "c": 100}
	fmt.Println(map1, map2, map3) // panic: assignment to entry in nil map

	map4 := make(map[string]int) // 不是nil
	fmt.Println(map4 == nil)
	map4["golang"] = 100
	fmt.Printf("%#v", map4)

	// 先创建一个map，然后在判断是否为nil
	var map5 map[string]int
	if map5 == nil {
		fmt.Println("map5 nil")
		map5 = make(map[string]int)
	}
	fmt.Println(map5 == nil)

}
