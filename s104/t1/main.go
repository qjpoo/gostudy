package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	/*  time
	1y = 365d
	1d = 24h
	1m = 60s
	1s = 1000ms   毫秒
	1ms = 1000us  微秒
	1us = 1000ns  ns纳秒
	1ns = 1000ps  ps皮秒


	*/
	// 获取当前时间及当天的其他时间点
	t1 := time.Now()
	fmt.Printf("%T, %v\n", t1, t1)
	today := time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), time.Local)
	fmt.Println(today)
	today1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(today1)

	// 获取其他日期,比如昨天
	yesterday := today1.AddDate(0, 0, -1)
	fmt.Println(yesterday)

	// 获取时间戳
	begin := yesterday.Unix()
	fmt.Println(begin)

	// 时间的格式化  //使用模板在对应时区转化为time.time类型  string ==> time
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-28 18:22:54", time.Local)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %v\n", tt, tt)

	tt1, err := time.Parse("2006-1-2 15:04:05", "2020-7-29 09:52:23")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %v\n", tt1, tt1)

	//获取本地location
	toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600

	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)

	// time.Unix()函数返回公元1970年1月，1日，0时，0分，0秒以来的秒数
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "1970-01-01 00:00:45", time.UTC)
	fmt.Printf("%T\n", t)
	fmt.Println(t)              // 1970-01-01 00:00:45 +0000 UTC
	fmt.Println("t:", t.Unix()) // t: 45

	// 本地时间1970年1月1日8时0分50秒,因为是中国时间是东八区,比UTC时间早8个小时
	// 相当于UTC的"1970-01-01 08:00:50",所以会打印出 t:50
	t0, _ := time.ParseInLocation("2006-01-02 15:04:05", "1970-01-01 08:00:50", time.Local)
	fmt.Println("t:", t0.Unix())

	fmt.Println("------------------------------------")
	//time.Time()声明时，若未初始化则表示UTC时间，公元1年，1月，1日，0时，0分，0秒
	t11 := time.Time{}
	t21 := time.Now()
	t1Second := t11.Unix()
	t2Second := t21.Unix()
	fmt.Printf("t1:%v, t1Second:%v \n", t11, t1Second) // UTC时间1970-01-01 08:00:50的秒数为0,所以在此时间之前的为负值,且在int32范围之外
	fmt.Printf("t2:%v, t2Second:%v \n", t21, t2Second)
	fmt.Println("t2Second-t1Second=", t2Second-t1Second) // 所以会返回一个超大的值,已超过int32范围

	//秒数转Duration，例如 dur := 2*time.Second + 500*time.Millisecond
	t10 := time.Now()                           // 当前时间
	time.Sleep(3 * time.Second)                 // 休眠3秒
	dur := 2*time.Second + 500*time.Millisecond // 规定间隔:2秒500毫秒
	tn := time.Now().Sub(t10)
	fmt.Printf("%T,%v\n", tn, tn) // time.Duration,3.0003995s
	// 判断间隔时
	if time.Now().Sub(t10) > dur {
		fmt.Println("sleep over 2.5s")
	}

	//time.Duration 转秒数，接示例3：durSec := dur.Seconds()，返回值类型为float64类型
	dur1 := 2*time.Second + 500*time.Millisecond
	durSec1 := dur1.Seconds()
	fmt.Println("durSec:", durSec1)

	//比较两个时间Time是否相等
	// 解析字符串 "2019-06-18 08:47:50"为本地时间
	t1a, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-06-18 08:47:50", time.Local)
	fmt.Printf("%T\n", t1a) // time.Time
	t2a, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-06-18 08:47:50", time.Local)
	if t1a.Equal(t2a) {
		fmt.Println("t1a == t2a")
	}

	t1b, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-06-18 08:47:50", time.Local)
	fmt.Printf("%T\n", t1b) // time.Time
	t2b, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-29 08:47:50", time.Local)
	if t1b.Before(t2b) {
		fmt.Println("t1b before t2b")
	}

	t1c, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-06-18 08:47:50", time.Local)
	fmt.Printf("%T\n", t1c) // time.Time
	t2c, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-06-18 08:47:50", time.Local)
	if t1c.After(t2c) {
		fmt.Println("t1c after t2c")
	}

	fmt.Println("-----------------------------------------")
	// 得到当年，当月，当日，0时，0分，0秒的时间字符串
	t1d := time.Now() // 2019年6月18日
	fmt.Printf("%T\n", t1d)
	curDayStr := t1d.Format("2006-01-02 00:00:00")   // 固定到天
	curMonthStr := t1d.Format("2006-01-00 00:00:00") // 固定到月
	fmt.Printf("curDayStr:%v \n", curDayStr)
	fmt.Printf("curMonthStr:%v \n", curMonthStr)

	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Nanosecond())

	fmt.Println("-----------------------------------------")
	fmt.Println(time.Now().Format("2006年/01月/02日 15:04:05"))
	fmt.Println(time.Now().Format(time.RFC3339))

	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Zone())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().Location())

	// 时间段(Duartion)
	fmt.Println("-----------------------------------------")
	d, _ := time.ParseDuration("100s")
	fmt.Printf("%T, %v\n", d, d)
	fmt.Println(d.Truncate(1000), d.Seconds(), d.Nanoseconds())

	// 时区(Location)
	local, _ := time.LoadLocation("")
	fmt.Println(local)

	// 默认UTC
	local1, _ := time.LoadLocation("")
	fmt.Println(local1)
	// 服务器设定的时区，一般为CST
	local2, _ := time.LoadLocation("Local")
	fmt.Println(local2)
	// 美国洛杉矶PDT
	local3, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(local3)

	// 获取指定时区的时间点
	local4, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Date(2018, 1, 1, 12, 0, 0, 0, local4))

	// 时间运算
	// func Sleep(d Duration)   休眠多少时间，休眠时处于阻塞状态，后续程序无法执行
	time.Sleep(time.Duration(10) * time.Second)

	// func After(d Duration) <-chan Time  非阻塞,可用于延迟
	time.After(time.Duration(10) * time.Second)

	// func Since(t Time) Duration 两个时间点的间隔
	start := time.Now()
	fmt.Println(time.Since(start)) // 等价于 Now().Sub(t)， 可用来计算一段业务的消耗时间

	//func Until(t Time) Duration     //  等价于 t.Sub(Now())，t与当前时间的间隔

	// func (t Time) Add(d Duration) Time
	//fmt.Println(dt.Add(time.Duration(10) * time.Second)) // 加

	//func (t Time) Sub(u Time) Duration                    // 减

	// func (t Time) AddDate(years int, months int, days int) Time
	//fmt.Println(dt.AddDate(1, 1, 1))

	// func (t Time) Before(u Time) bool
	// func (t Time) After(u Time) bool
	// func (t Time) Equal(u Time) bool          比较时间点时尽量使用Equal函数

	// 日期时间差
	dt1 := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	dt2 := time.Date(2018, 1, 9, 23, 59, 22, 100, time.Local)
	// 不用关注时区，go会转换成时间戳进行计算
	fmt.Println(dt1.Sub(dt2))

	//基于当前时间的前后运算
	now := time.Now()

	// 一年零一个月一天之后
	fmt.Println(now.Date())

	// 一段时间之后
	fmt.Println(now.Add(time.Duration(10) * time.Minute))

	// 计算两个时间点的相差天数
	dt1 = time.Date(dt1.Year(), dt1.Month(), dt1.Day(), 0, 0, 0, 0, time.Local)
	dt2 = time.Date(dt2.Year(), dt2.Month(), dt2.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(int(math.Ceil(dt1.Sub(dt2).Hours() / 24)))

	// 时区转换
	// time.Local 用来表示当前服务器时区
	// 自定义地区时间
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Println(time.Date(2018, 1, 2, 0, 0, 0, 0, beijing)) // 2018-01-02 00:00:00 +0800 Beijing Time

	// 当前时间转为指定时区时间
	fmt.Println(time.Now().In(beijing))

	// 指定时间转换成指定时区对应的时间
	dt, err := time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local)

	// 当前时间在零时区年月日   时分秒  时区
	year1, mon1, day1 := time.Now().UTC().Date()  // 2018 April 24
	hour1, min1, sec1 := time.Now().UTC().Clock() // 3 47 15
	zone1, _ := time.Now().UTC().Zone()         // UTC
	fmt.Println(year1, mon1, day1, hour1,min1, sec1, zone1)

	//比较两个时间点
	dt0 := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	fmt.Println(time.Now().After(dt0))  // true
	fmt.Println(time.Now().Before(dt0)) // false

	// 是否相等 判断两个时间点是否相等时推荐使用 Equal 函数
	fmt.Println(dt.Equal(time.Now()))

	// 设置执行时间
	//通过time.After 函数与 select 结合使用可用于处理程序超时设定

/*	select {
	case m := <-c:
		// do something
	case <-time.After(time.Duration(1) * time.Second):
		fmt.Println("time out")
	}
*/
	//Ticker类型
	//Ticker 类型包含一个 channel，有时我们会遇到每隔一段时间执行的业务(比如设置心跳时间等)，就可以用它来处理，这是一个重复的过程

/*	// 无法取消
	tick := time.Tick(1 * time.Minute)
	for _ = range tick {
		// do something
	}

	// 可通过调用ticker.Stop取消
	ticker := time.NewTicker(1 * time.Minute)
	for _ = range tick {
		// do something
	}
*/
}
