package main

import (
	"fmt"
	"reflect"
	"time"
)

const LAYOUT = "2016年01月02日 15:04:05"
var shanghaiTimeZone *time.Location

func initTimeZone()  {
	if location, err := time.LoadLocation("Asia/Shanghai");err !=nil {
		fmt.Println("failed to load timezone, err: ", err)
	}else {
		shanghaiTimeZone = location
	}
}
func main() {
	now := time.Now()
	fmt.Println(reflect.TypeOf(now)) // time.Time
	nowstr := now.Format(LAYOUT)
	fmt.Println(nowstr)  //
	fmt.Println(reflect.TypeOf(nowstr))  // string

	initTimeZone()
	now1 := time.Now().In(shanghaiTimeZone)
	nowstr1 := now1.Format(LAYOUT)
	fmt.Println(nowstr1)

	//
	fmt.Println("---------------------------------------")
	var t time.Time // 定义 time.Time 类型变量
	t = time.Now()  // 获取当前时间
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t, t.Location(), t)
	// 时间: 2017-02-22 09:06:05.816187261 +0800 CST, 时区:  Local,  时间类型: time.Time

	// time.UTC() time 返回UTC 时区的时间
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t.UTC(), t.UTC().Location(), t)
	// 时间: 2017-02-22 01:07:15.179280004 +0000 UTC, 时区:  UTC,  时间类型: time.Time

	// 获取一个时间的方法
	fmt.Println("---------------------------------------")
	// 当前本地时间
	t = time.Now()
	fmt.Println("'time.Now': ", t)

	// 根据时间戳返回本地时间
	t_by_unix := time.Unix(1596013056, 0)
	fmt.Println("'time.Unix': ", t_by_unix)

	// 返回指定时间
	t_by_date := time.Date(2017, time.Month(2), 23, 1, 30, 30, 0, time.Local)
	fmt.Println("'time.Date': ", t_by_date)
	fmt.Println(time.Now().Unix())

	// 时间显示
	// 获取指定时间在UTC 时区的时间表示
	fmt.Println("---------------------------------------")
	t_by_utc := t.UTC()
	fmt.Println("'t.UTC': ", t_by_utc)

	// 获取本地时间表示
	t_by_local := t.Local()
	fmt.Println("'t.Local': ", t_by_local)

	// 时间在指定时区的表示
	t_in := t.In(time.UTC)
	fmt.Println("'t.In': ", t_in)

	// Format
	fmt.Println("t.Format", t.Format(time.RFC3339))


	fmt.Println("---------------------------------------")
	// 时间序列化
	// 时间序列化
	t_byte, err := t.MarshalJSON()
	fmt.Println("'t.MarshalJSON': ", t_byte, err)

	// 时间数据反序列化
	var t_un time.Time
	err = t_un.UnmarshalJSON(t_byte)
	fmt.Println("'t_un.UnmarshalJSON': ", t_un, err)


}
