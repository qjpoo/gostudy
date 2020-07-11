package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		求素数
		打印2-100内的素数(只能被1和本身整除)
	2 3 5 7 11 13 17
	*/

	for i := 2; i <= 100; i++ {
		flag := true
		//for j := 2; j < i; j++ {
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i % j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
