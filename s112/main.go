package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 遍历文件夹


	// fmt.Println(os.Getwd())
	fileDir := "C:\\Users\\瞿健\\go\\src\\awesomeProject"
	listFiles(fileDir, 0)

}
func listFiles(dirname string, level int)  {
	s := "|--"
	for i:=0;i<level;i++ {
		s = "|  " + s
	}
	fileInfo, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, fi := range fileInfo {
		fileName := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, fileName)
		if fi.IsDir() {
			listFiles(fileName, level+1)
		}
	}
}
