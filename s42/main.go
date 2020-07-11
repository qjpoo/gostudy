package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		随机数
		math.rand
	*/

	num1 := rand.Int()
	fmt.Println(num1)
	fmt.Println("==============================")

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}

	fmt.Println("==============================")
	rand.Seed(100)
	fmt.Println(rand.Intn(10))

	t := time.Now()
	fmt.Println(t)        // 2020-07-11 20:40:03.5323639 +0800 CST m=+0.002000101
	fmt.Printf("%T\n", t) // time.Timeint64

	ts := t.Unix()  // 秒
	fmt.Println(ts) //
	fmt.Printf("%T\n", ts)

	ts1 := t.UnixNano()
	fmt.Println(ts1) //
	fmt.Printf("%T\n", ts1)

	// 生成随机数函数
	// step1 生成种子
	rand.Seed(time.Now().UnixNano())
	//step2 生成随机数
	for i := 1; i <= 10; i++ {
		r := rand.Intn(100)
		fmt.Println(r)
	}

	// 范围随机数
	/*
		Intn(n) ==> 0 - n-1
		要获得3-48之间的数    rand.Intn(46) + 3
	    [15,76]  rand.Int(62) + 15
	*/

	fmt.Println(rand.Intn(46) + 3)
}
