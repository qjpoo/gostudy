package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Hostname())
	fmt.Println(os.UserHomeDir())

	s1 := os.Environ()
	for _, v := range s1 {
		fmt.Println(v)
	}

	fmt.Println(os.Getenv("windir"))

	fmt.Println(os.Getpid())

	// os.Exit(0)


	fmt.Println("-------------------------")
	f, err := os.Stat("C:\\Users\\86137\\.viminfo")
	if err !=nil {
		fmt.Println(err)
	}else {
		fmt.Println(f.Name())
		fmt.Println(f.IsDir())
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(pwd)
	}

	err1 := os.Mkdir("c:\\testdir", 0755)
	fmt.Println(err1)

	err2 := os.Remove("c:\\testdir")
	fmt.Println(err2)

	tmpDir := os.TempDir()
	fmt.Println(tmpDir)


	fmt.Println("-------------------------")
	dir, err := os.Getwd()
	fmt.Println(dir, err)
	file := dir + "/new.txt"

	var fh *os.File
	fi, _ := os.Stat(file)
	//fmt.Println(fi)
	if fi == nil {
		fh, _ = os.Create(file)
	}else {
		fh, _ = os.OpenFile(file, os.O_RDONLY, 0666)
	}
	fmt.Printf("%T, %v", fh, fh)

	w := []byte("hello go language " + time.Now().String())
	n, err := fh.Write(w)
	fmt.Println(n, err)

	ret, err := fh.Seek(0,0)
	fmt.Println("当前文件的指针位置： ", ret, err)

	b := make([]byte, 128)
	n, err = fh.Read(b)
	fmt.Println(n, err, string(b))

	fh.Close()



}
