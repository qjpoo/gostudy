package main

import "fmt"

func main() {
	/*
		三位数: [100, 999]
		水仙花数: 每个位上的数字立方和,刚好等 该数字本身
		比如: 153 = 1 * 1 * 1 + 5 * 5 * 5 + 3 * 3 *
	*/
	g := 0
	s := 0
	b := 0
	for i := 100; i <= 999; i++ {
		b = i / 100
		s = (i - b*100) / 10
		g = (i - b*100 - s*10)
		//fmt.Println(b, s, g)
		var x int
		x = b*b*b + s*s*s + g*g*g
		if i == x {
			fmt.Println(i)
		}
	}

	// 逆向思维,凑数
	/*
		百位是: 1 - 9
		十位是: 0 - 9
		个位是: 0 - 9
	*/
	fmt.Println("----------------------------------")
	for x := 1; x <= 9; x++ {
		for y := 0; y <= 9; y++ {
			for z := 0; z <= 9; z++ {
				m := x*100 + y*10 + z
				if m == x*x*x+y*y*y+z*z*z {
					fmt.Println(m)
				}
			}
		}
	}

}
