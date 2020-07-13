package main

import "fmt"

func main() {
	urls := make(map[string]string, 3)
	//var urls = map[string]string{}
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"
	names := make([]string, 0)
	for key, _ := range urls {
		names = append(names, key)
	}
	fmt.Println(names, len(names))
	fmt.Printf("%T, %T", urls, names)
}
