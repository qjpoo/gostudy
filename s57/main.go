package main

import "fmt"

func main() {
	/*
	map结合slice的使用
	1. 创建map用于存储人的信息 name, age, sex, address
	2. 每个map存储一个人的信息
	3. 将这样map存入到slice中
	*/

	map1 := make(map[string]string)
	map1["name"] = "qujian"
	map1["age"] = "20"
	map1["sex"] = "男"
	map1["address"] = "湖北洪湖"

	map2 := make(map[string]string)
	map2["name"] = "chiling"
	map2["age"] = "30"
	map2["sex"] = "男"
	map2["address"] = "湖南岳阳"

	map3 := map[string]string{"name": "xi", "age": "28", "sex":"女", "address": "北京"}
	fmt.Println(map1, map2,map3)

	// 将map放入到slice里面
	s := make([]map[string]string, 0)
	s = append(s, map1)
	s = append(s, map2)
	s = append(s, map3)

	for k, v := range s {
		fmt.Printf("第 %d 个人的信息: ", k)
		fmt.Printf("name: %s\n", v["name"])
		fmt.Printf("age: %s\n", v["age"])
		fmt.Printf("sex: %s\n", v["sex"])
		fmt.Printf("address: %s\n", v["address"])
	}


}
