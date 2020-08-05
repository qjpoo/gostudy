package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2, 5,3, 6,1, 10}
	sort.Ints(s)
	fmt.Println(s)

	s2 := []int{2, 5,3, 6,1, 10}
	sort.Sort(sort.Reverse(sort.IntSlice(s2)))
	fmt.Println(s2)

	str := []string{"c", "a", "b", "A", "I"}
	sort.Strings(str)
	fmt.Println(str)

	str1 := []string{"c", "a", "b", "A", "I"}
	sort.Sort(sort.Reverse(sort.StringSlice(str1)))
	fmt.Println(str)

	
}
