package main

import "fmt"

func main() {

	// map是引用, map, slice , chan
	// 用make创建的都是引用类型

	map1 := make(map[string]map[string]string)
	m1 := make(map[string]string)
	m1["name"] = "chilig"
	m1["age"] = "20"
	m1["addr"] = "湖北"
	map1["hr"] = m1

	m2 := make(map[string]string)
	m2["name"] = "qujian"
	m2["age"] = "21"
	m2["addr"] = "湖南"
	map1["coo"] = m2

	fmt.Println(map1) // map[coo:map[addr:湖南 age:21 name:qujian] hr:map[addr:湖北 age:20 name:chilig]]

	ma1 := make(map[string]string)
	ma1["name"] = "qujian"
	ma1["age"] = "21"
	ma1["addr"] = "湖南"
	ma1["coo"] = "xxoo"
	fmt.Println(ma1)

	ma2 := ma1
	ma2["coo"] = "xo"
	fmt.Println(ma2, ma1)

}
