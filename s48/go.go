package main

import "fmt"

func main() {

	/*
	二维数组
	*/
	a := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a)
	fmt.Println(&a[0])
	fmt.Println(len(a)) // 二维数组的长度
	fmt.Println(len(a[0])) //一维数组的长度
	fmt.Println(a[1][2]) //一维数组的长度
	fmt.Println("=====================")

	for i:=0; i<len(a);i++ {
		for j:=0; j<len(a[i]);j++ {
			//fmt.Printf("%v\t",a[i][j])
			fmt.Print(a[i][j],"\t")
		}
		fmt.Println()
	}

	fmt.Println("=====================")
	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, "\t")
		}
		fmt.Println()
	}
}
