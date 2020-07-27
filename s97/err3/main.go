package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	files, err := filepath.Glob("[")
	fmt.Println(files, err)
	if err != nil && err == filepath.ErrBadPattern{
		fmt.Println(err)
		return
	}
	fmt.Println(files)
}
