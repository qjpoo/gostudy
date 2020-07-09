package main

import "fmt"

func main() {
	// iota

	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Println(a, b, c)

	const (
		a1 = iota // 0
		b1 // b1 = iota 1
	)
	fmt.Println(a1, b1)

	const (
		A = iota // 0
		B // 1
		C // 2
		D = "hi" // 3
		E // E = "hi" 4
		F = 100 // 5
		G // G = 6
		H = iota // H =  7
		I  // I = 8
	)

	const (
		J = iota  // 0
	)
	fmt.Println(A,B,C,D,E,F,G,H,I,J)

}