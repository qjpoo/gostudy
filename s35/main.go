package main

import "fmt"

func main() {
	// 打印58-23的数字
	for i := 58; i >= 23; i-- {
		fmt.Println(i)
	}

	// 100以内求和
	sum := 0
	for i :=1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 打印1-100内,能被3整除,不能被5整除的数字,统计的数字每行打印5个
	count := 0
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 != 0{
				fmt.Printf("%d\t",i)
				count++
				if count % 5 == 0 {
					fmt.Println()
				}

			}
		}
}
