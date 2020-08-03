package main

import (
	"bytes"
	"fmt"
)

func main() {

	var b bytes.Buffer
	_, err := b.Write([]byte("hello, world ..."))
	fmt.Println(b, err)


}
