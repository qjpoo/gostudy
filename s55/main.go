package main

import "fmt"

func main() {
	// map的使用
// nil的map是不能使用的


	// map的定义和存储值
	map1 := make(map[int]string)
	map1[0] = "Golang"
	map1[1] = "Java"
	map1[2] = "C"
	map1[3] = "Python"
	map1[4] = "Ruby"
	map1[5] = ""
	fmt.Printf("%T， %#v\n", map1, map1)
	fmt.Println(map1[0])
	fmt.Println(map1[5]) // 获取的是一个空的串，如果是一个int，那就是0

	// 获得的是不是一个Key值是否存在, 值为空值
	v1, ok := map1[5]
	if ok {
		fmt.Printf("值为：%#v\n", v1)
	}else {
		fmt.Println("空值")
	}

	// 不存在的值
	v2, ok := map1[6]
	if ok {
		fmt.Printf("值为：%#v\n", v2)
	}else {
		fmt.Println("空值")
	}

	// 修改ma的值
	fmt.Println(map1)
	map1[0] = "C++"
	fmt.Println(map1)

	// 删除
	delete(map1,0)
	fmt.Println(map1)
	//删除一个不存在的Key
	delete(map1, 60)

	//长度
	fmt.Println(len(map1))
}
