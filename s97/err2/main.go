package main

import (
	"fmt"
	"net"
)

func main() {
	//


	//add, err := net.LookupHost("www.baidu.com")
	add, err := net.LookupHost("www.baidu12323232.com")
	fmt.Println(add) // [14.215.177.38 14.215.177.39]
	//fmt.Println(err)
	if ins, ok := err.(*net.DNSError);ok {
		if ins.Timeout(){
			fmt.Println("timeout ...")
		}else if ins.Temporary() {
			fmt.Println("临时性的错误...")
		}else {
			fmt.Println("error ...")
		}
	}
}
