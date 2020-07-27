package main

import (
	"fmt"
	"os"
)

func main() {
	// 错误的类型表示

	f, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("err open file failure ...")
		fmt.Println(err)
		if ins, ok := err.(*os.PathError);ok {
			fmt.Println("1. Op", ins.Op)
			fmt.Println("2. Path", ins.Path)
			fmt.Println("3. Err", ins.Err)
		}
		return
	}
	fmt.Println("open file sucessfull ...", f.Name())
	defer f.Close()
}
