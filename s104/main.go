package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	/*
	time包
	1s = 1000ms
	1ms = 1000us
	1us = 1000ns
	1ns = 1000ps 皮秒



	*/

	//1. 获取当前的时间
	t1 := time.Now()
	fmt.Println(reflect.TypeOf(t1), t1) // time.Time

	// 2. 获取指定的时间
	t2 := time.Date(2020,07,29,23,35,10,20,time.Local)
	fmt.Println(t2)

	// 3. time ==> string
	// t.format 里面参数是layout模版
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)
	s2 := t1.Format("2006年1月2日")
	fmt.Println(s2)

	// string ==> time
	// time.parse("layout", string)
	s3 := "1999年9月9号"
	t3, err := time.Parse("2006年1月2号",s3)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(reflect.TypeOf(t3), t3)
	}

	fmt.Println("--------------------------")
	fmt.Println(reflect.TypeOf(t1), t1) // time.time
	fmt.Println(reflect.TypeOf(t1.String()), t1.String()) // string

	// 4. 根据当前的时间,获取指定的内容
	y, m, d := t1.Date() // year month day
	fmt.Println(y, m, d)
	h, min, s := t1.Clock()  // hour min, sec
	fmt.Println(h, min, s)
	fmt.Println(t1.YearDay()) // 211 今年一共过了多少天
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())
	fmt.Println(t1.Hour())
	fmt.Println(t1.Minute())
	fmt.Println(t1.Second())
	fmt.Println(t1.Nanosecond())
	fmt.Println(t1.Weekday())

	// 5. 时间戳: 距离1970.01.01 00:00:00到指定时间的秒, 纳秒
	t4 := time.Date(1970,01,01,00,01,0,0,time.UTC) // 注意这里是UTC,不是time.local
	timeStamp := t4.Unix() // 秒的差值
	fmt.Println(timeStamp)
	// 求一下当前的时间到1970过去的秒数
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Nanosecond())

	// 6. 时间间隔
	t5 := t1.Add(time.Minute)
	fmt.Println(t1, t5)
	fmt.Println(t1.Add(time.Hour * 12))

	t6 := t1.AddDate(1,0,0) // 加一年
	fmt.Println(t6)

	t61 := t1.AddDate(0,-1,0) // 减一个月
	fmt.Println(t61)

	d1 := t1.Sub(t6)
	fmt.Println(d1)

	// 7. 睡觉
	time.Sleep(3 * time.Second) // 睡觉3秒

	// 睡眠1-10随机数
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1 // int
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum) * time.Second)
	fmt.Println("wake up ...")

}
