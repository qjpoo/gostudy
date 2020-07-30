package main

import (
	"fmt"
	"strings"
)

func main() {

	s, t := "hello go", "hello Go"
	isEqual := strings.EqualFold(s, t)
	fmt.Println(isEqual)

	//prefix := "/root/123.txt"
	prefix := "Abc.txt"
	//isTrue := strings.HasPrefix(prefix, "/")
	isTrue := strings.HasPrefix(prefix, "a")
	fmt.Println(isTrue)

	suffix := "/root/123.txt"
	isTrue1 := strings.HasSuffix(suffix, ".txt")
	fmt.Println(isTrue1)

	subStr := "txt"
	isTrue2 := strings.Contains(suffix, subStr)
	fmt.Println(isTrue2)

	r := rune(101)
	fmt.Println(string(r))
	conRun := strings.ContainsRune(s, r)
	fmt.Println(conRun)

	sep := "0"
	sepIndex := strings.Index(s, sep)
	fmt.Println(sepIndex)







}
