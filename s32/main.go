package main

import "fmt"

func main() {
	/*
	break fallthrough, 这两个效果是相反的,一个跳出,一个是继续
	break 强字结束break语句
	fallthrough 穿透,不用匹配后面的case的条件了
	*/

	n := 2
	switch n {
	case 1:
		fmt.Println("11111111111111111")
		fmt.Println("11111111111111111")
		fmt.Println("11111111111111111")
		fmt.Println("11111111111111111")
	case 2:
		fmt.Println("2222222222")
		break //swithc结束了
		fmt.Println("2222222222")
		fmt.Println("2222222222")
		fmt.Println("2222222222")
	case 3:
		fmt.Println("3333333333333")
		fmt.Println("3333333333333")
		fmt.Println("3333333333333")
		fmt.Println("3333333333333")
	default:
		fmt.Println("default")
	}
	fmt.Println("main over ...")


	// fallthrough
	i := 4
	switch i {
	case 1:
		fmt.Println("1111111111111111")
	case 2:
		fmt.Println("222222222222222")
	case 3:
		fmt.Println("333333333333333")
	case 4:
		fmt.Println("444444444444444444")
		fallthrough //fallthrough应该是在case这个语句块的最后一行,它的后面不能在有语句了
	default:
		fmt.Println("default") // 不用判断条件了,直接打应里面的语句块
	}
}
