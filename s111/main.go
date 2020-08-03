package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	// ioutil io操作的辅助包

	// ReadFile 读取文件中所有的数据
	fileName := "123.txt"
	data, err := ioutil.ReadFile(fileName)
	fmt.Println(string(data), err)

	// WriteFile 写之前会,清空里面的数据
	fileName = "n.txt"
	err = ioutil.WriteFile(fileName, []byte("hello, xoo"), 0777)
	if err != nil {
		fmt.Println(err)
	}

	// ReadAll, ReadFile底层也是调用的ReadAll
	s2 := "xxoo, this is a teacher"
	r1 := strings.NewReader(s2)
	//s := make([]byte, 1024, 1024)
	//xx, err := r1.Read(s)
	//fmt.Println(xx, string(s), err)
	data, err = ioutil.ReadAll(r1)
	fmt.Println(string(data), err)

	// ReadDir, 只能读取一层
	//fmt.Println(os.Getwd())
	dirName := "C:\\Users\\瞿健\\go\\src\\awesomeProject"
	fileInfo, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(fileInfo))
	for i := 0; i < len(fileInfo); i++ {
		fmt.Println(i, fileInfo[i].Name(), fileInfo[i].Size(), fileInfo[i].IsDir())
	}

	// Tempdir
	str, err := ioutil.TempDir("C:\\Users\\瞿健\\go\\src\\awesomeProject", "tmpdir")
	fmt.Println(str, err) // tmpdir831243619
	defer os.Remove(str)

	// TempFile
	f1, err := ioutil.TempFile("C:\\Users\\瞿健\\go\\src\\awesomeProject", "tmpfile.tmp")
	fmt.Println(f1.Name(), err) // tmpfile.tmp084413798
	defer os.Remove(f1.Name())

	s := "golang是世界上最好的语言"
	cnt := 0
	for _, v := range s {
		//fmt.Println(reflect.TypeOf(v))
		//fmt.Println(string(v))
		if unicode.Is(unicode.Han, v) {
			cnt++
		}
	}
	fmt.Println(cnt)

}
