package main

import "fmt"

// 1.定义接口
type USB interface {
	start() // USB的开始工作   只写方法的定义
	end() // 结束工作  只写方法的定义
}

// 2. 实现类
type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

// 3. 让类实现接口中的方法
func (m Mouse) start()  {
	fmt.Println(m.name, "start ...")
}

func (m Mouse) end()  {
	fmt.Println(m.name, "end  ...")

}


func (f FlashDisk) start()  {
	fmt.Println(f.name, "start ...")
}

func (f FlashDisk) end()  {
	fmt.Println(f.name, "end  ...")

}

// 测试方法
func testUsbInter(usb USB)  {
	usb.start()
	usb.end()
}

func (f FlashDisk) deleteData()  {
	fmt.Println(f.name, "删除数据")
}
func main() {
	/*
	接口: 一组方法签名
	定义接口,就是定义一组方法
	接口中只是定义一些方法,但是没有定义方法的实现,谁实现了接口中的方法,谁就实现了该接口

	go语言中,接口和类型的实现关系,是非侵入式.
	就是说不用说明那个类型实现了这个接口,比如Mouse实现了usb接口
	而别的像Java语言要说说明: class Mouse implements USB()



	 */
	// 1. 当需要接口类型的对象时，可以使用任意实现类对象代替
	m1 := Mouse{name: "罗技"}
	fmt.Println(m1.name)

	f1 := FlashDisk{name: "金士顿"}
	fmt.Println(f1.name)
	f1.deleteData()


	testUsbInter(m1)  // usb = m1
	testUsbInter(f1)  // usb = f1

	//
	var usb USB
	usb = m1
	// fmt.Println(usb.name)  // 不能访问name字段，因为usb里面没有定义name
	usb.start()
	usb.end()

	usb = f1
	// usb.deleteData()　usb是访问不了f1里面的deleteData()方法

	var arr [3]USB
	arr[0] = f1
	arr[1] = m1
	fmt.Println(arr)

}
