package main

import "fmt"

func sum(x int) int {
	s := 0
	for i:=1;i<=x;i++{
		s += i
	}
	return s
}
func main() {
	// function
	fmt.Println(sum(100))

}
