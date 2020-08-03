package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

func main() {
	// path package

	// 判断是否是一个绝对路径
	pt := "/root/home/x1/test/hello.sh"
	is1 := path.IsAbs(pt)
	fmt.Println(is1)

	// 将路径分割为路径和文件名
	dir, file := path.Split(pt)
	fmt.Println(dir, file) // /root/home/x1/ test

	// 将多个字符串合并为一个路径
	jo := path.Join("/usr", "local", "src")
	fmt.Println(jo)

	// 返回路径中扩展
	fmt.Println(path.Ext(pt)) // .sh

	// 返回路径的最后一个元素
	fmt.Println(path.Base(pt)) // hello.sh

	// 返回目录
	fmt.Println(path.Dir(pt)) // /root/home/x1/test

	// 返回目录中的最短路径
	dira := "/usr/../hello/../local/src/ssh"
	fmt.Println(path.Clean(dira)) // /local/src/ssh

	// 正则是否匹配路径
	isMatch, err := path.Match(".*.sh", "/root/hello.sh")
	fmt.Println(isMatch, err)

	/*var arr = [...]int{5:2}
	fmt.Println(arr)
	var pf *[6]int = &arr
	fmt.Println((*pf)[5])*/

	// ------------------------------------------------------
	// path/filepath 包
	b1 := filepath.Base("/root/home/file.go")
	fmt.Println(b1)
	sep := filepath.Separator
	fmt.Println(reflect.TypeOf(sep), string(sep)) // int32 run
	fmt.Println(string(101))
	// 预定义变量  预定义变量，表示路径分隔符 /  预定义变量，表示环境变量分隔符 :
	// \ ; \ ;
	fmt.Println(string(filepath.Separator), string(filepath.ListSeparator), string(os.PathSeparator), string(os.PathListSeparator))

	// 返回绝对路径
	dir1 := "C:\\Users\\86137\\godfd\\df"
	real_dir, err := filepath.Abs(dir1)
	fmt.Println(real_dir, err)

	// 返回targpath 相对 basepath路径
	basePath, targetPath := "/usr/local", "/usr/local/go/bin"
	relDir, err := filepath.Rel(basePath, targetPath)
	fmt.Println(relDir, err) // go\bin

	fmt.Println("------------------------")
	// 返回路径最前面的卷名
	root := "/usr/local/go"
	vol := filepath.VolumeName(root)
	fmt.Println(vol) // ''

	// 路径分隔符替换为 `/`
	slash_dir := filepath.ToSlash(root)
	fmt.Println(slash_dir) // /usr/local/go

	//  `/` 替换为路径分隔符
	from_slash := filepath.FromSlash(slash_dir)
	fmt.Println(from_slash) // /usr/local/go



	targetPath1 := os.Getenv("path")
	pathSlice1 := filepath.SplitList(targetPath1)
	for _, v := range pathSlice1 {
		fmt.Println(v)
	}

	// 遍历当前目录的文件树
	rootDir, err := os.Getwd()
	err = filepath.Walk(rootDir, pathName)
	fmt.Println(err)

}
func pathName(path string, info os.FileInfo, err error) error  {
	if err != nil {
		return err
	}else {
		fmt.Println(path)
		fmt.Println(info.Size())
	}
	return nil

}
