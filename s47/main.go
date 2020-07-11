package main

import "fmt"

func main() {
	/*
		数组的排序:
			让数组的元素有一点的顺序
		var arr := [5]int{15, 23, 8, 10, 7}
		升序和降序
	*/

	/*
	第一轮
		for i:=0;i<4;i++ {
				if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}

			第二轮
				for i:=0;i<3;i++ {
						if arr[i] > arr[i+1] {
						arr[i], arr[i+1] = arr[i+1], arr[i]
					}
			}
	*/
	var arr = [5]int{15, 23, 8, 10, 7}
	for j := 1; j < len(arr); j++ {
		for i := 0; i < len(arr)-j; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		fmt.Println(arr)
	}
	fmt.Println("=================")
	fmt.Println(arr)

}
