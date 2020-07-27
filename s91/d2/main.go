package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%T\n", reader)
	fmt.Print("Please input something: ")
	//text, _ := reader.ReadString('\n')
	text, _ := reader.ReadBytes('\n')
	fmt.Printf("%T\n", text) // php think task stop
	fmt.Println(string(text))

}
