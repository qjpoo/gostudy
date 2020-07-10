package main

import "fmt"

func main() {
	/*
	字符串： 多个byte的集合
	语法： 使用双引号 "abc" "A" `xxoo`
	*/
	var s1 string
	s1 = "chiling"
	fmt.Printf("%T, %s\n", s1, s1)

	s2 := `hello world ...`
	fmt.Printf("%T, %s\n", s2, s2)

	// 区别 'A', "A"
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T, %d\n", v1, v1) // int32, 65
	fmt.Printf("%T, %s\n", v2, v2) // string, A

	v3 := '瞿'
	fmt.Printf("%T, %d\n", v3, v3) // int32, 30655

	v4 := '健'
	fmt.Printf("%T, %d\n", v4, v4) // int32, 20581
	fmt.Printf("%T, %d, %c, %q, %v\n", v4, v4, v4, v4, v4) // int32, 20581

	// 转义字符
	fmt.Println("'hello, \"world ...")

	// \n \t
	fmt.Println("\nhello\tworld ...")

	fmt.Println(`helle, "world"`)

}
