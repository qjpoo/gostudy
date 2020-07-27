package main

import "fmt"

type A interface {
}

type Cat struct {
	color string
}

type Person struct {
	name string
	age  int
}

// 接口A是空接口, 理解为可以代表任意的类型
func testA(a A) {
	fmt.Println(a)

}

func testB(a interface{}) {
	fmt.Println(a)
}

func sliceData(s []interface{}) {
	for i := 1; i <= len(s); i++ {
		fmt.Printf("第%d个元素是%v\n", i, s[i-1])
	}

}
func main() {
	/*
		空接口
			不包含任何的方法,正因为如此,所有的类型都实现了空接口,因此空接口可以存储任意的类型的数值
	*/

	var a1 A = Cat{color: "black"}
	var a2 A = Person{name: "quinn", age: 30}
	var a3 string = "hello"
	var a4 int = 10
	fmt.Printf("%T", a1)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	testA("hello, world ...")
	testB(10)

	// map key 是string value是任意类型
	map1 := make(map[string]interface{})
	map1["1"] = 1
	map1["2"] = 3.14
	map1["3"] = "hello"
	map1["4"] = Person{name: "xxoo", age: 18}
	fmt.Println(map1)

	// 切片, 存储任意类型的数据
	s1 := make([]interface{}, 0, 10)
	s1 = append(s1, a1, a2, a3, a4, 100, 3.14, "hello")
	fmt.Println(s1)

	sliceData(s1)

}
