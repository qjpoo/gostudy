package main

import "fmt"

func main() {
	var x []int
	x = make([]int, 3)
	fmt.Printf("%T", x)

	fmt.Println("=========================")
	y := []int{1, 2, 3, 4, 5}
	var z []int
	z=y
	fmt.Println(y,"\t", z)
	z = append(z, 6, 7, 8, 9, 10)
	fmt.Println(y,"\t", z)
	n1 := []string{"A", "B", "C"}
	n2 := []string{"A1", "B1", "C1"}
	copy(n1, n2) //把源y， 拷贝到z中
	fmt.Println(n1)
	n1 = append(n1, n2)
	fmt.Println(n1)

	fmt.Println("=========================")
	var m []int
	m = y[1:4]
	fmt.Println(m)

	x1 := [...]int{1, 2, 3, 4, 5, 6,7}
	x2 := x1[2:5]
	fmt.Println(x1, x2)
	for _, v := range x2 {
		fmt.Println(v)
	}

}
