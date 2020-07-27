package main

import (
	"fmt"
	"math"
)

// 定义一个接口
type Shape interface {
	peri() float64 // 周长
	area() float64 // 面积
}

// 定义一个三角形的类
type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

// 定义一个圆的类
type Cricle struct {
	r float64
}

// 面的周长
func (c Cricle) peri() float64 {
	return 2 * c.r * math.Pi
}

// 圆的面积
func (c Cricle) area() float64 {
	return math.Pow(c.r, 2) * math.Pi
}

func testShape(s Shape) { // 函数接受的是一个接口类型
	fmt.Printf("周长为: %f, 面积为: %f\n", s.peri(), s.area())

}

func getType(s Shape) {
	if ins, ok := s.(Triangle); ok {
		fmt.Println("三角形", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Cricle); ok {
		fmt.Println("圆形", ins.r)
	} else if ins, ok := s.(*Triangle); ok {
		/*
		t2: *main.Triangle, 0xc000006030, 0xc00000a500
		ins: *main.Triangle, 0xc000006038, 0xc00000a500
		s: *main.Triangle, 0xc00002a220, 0xc00000a500
		这里t2, ins, s的地址都不相同,但是它产里面存储的结构体是一样的,所以它们的地址是一样的,因为是指针传递
		*/
		fmt.Printf("ins: %T, %p, %p\n", ins, &ins, ins)
		fmt.Printf("s: %T, %p, %p\n", s, &s, s)
	} else {
		fmt.Println("未知形状 ...")
	}
}

func getType2(s Shape)  {
	switch ins := s.(type) {
	case Triangle:
		fmt.Printf("三角形三边为: %d, %d, %d\n", ins.a, ins.b, ins.c)
	case Cricle:
		fmt.Printf("圆形的半径为: %f", ins.r)
	case *Triangle:
		fmt.Printf("三角形指针, 三边为: %f, %f, %f", ins.a, ins.b, ins.c)
	}
	
}
func main() {
	/*
		接口断言:
		1. instance := 接口对象.(实际类型) // 不安全,会panic()
		2. instance, ok := 接口对象.(实际类型)  /// 安全

		方式二:
		switch instance := 接口对象.(type) {
		case 实际类型1:
			...
		case 实际类型2:
		    ...
		}
	*/

	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.a, t1.b, t1.c)
	fmt.Println(t1.peri())
	fmt.Println(t1.area())

	var c1 Cricle = Cricle{10}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())

	// 定义一个接口类型
	var s1 Shape
	s1 = t1 // 把t1赋值给s1 虽然说s1是一个三角形,但是它没有三角形的三边的属性
	fmt.Println(s1.peri())
	fmt.Println(s1.area())
	// fmt.Println(s1.a, s1.b, s1.c)  // 此时s1是没有三条边的属性的

	var s2 Shape
	s2 = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())
	// fmt.Println(s2.r) 没有圆的半径属性

	testShape(c1)
	testShape(t1)
	testShape(s2)
	testShape(s1)

	fmt.Println("----------------")
	getType(t1)
	getType(c1)
	getType(s1)

	fmt.Println("----------------")
	var t2 *Triangle = &Triangle{3, 4, 2}
	fmt.Printf("t2: %T, %p, %p\n", t2, &t2, t2)
	getType(t2)


	getType2(t2)
}
