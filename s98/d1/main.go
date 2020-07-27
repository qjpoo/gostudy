package main

import (
	"fmt"
	"math"
)

// 1. 定义一个结构体，表示错误的类型
type areaError struct {
	msg string
	radius float64
}

// 2. 实现error接口，就是实现error()访求
func (e *areaError) Error() string  {
	return  fmt.Sprintf("error: 半径: %.2f, %s", e.radius, e.msg)
}

// 圆形的面积
func roundArea(radius float64) (float64, error)  {
	if radius < 0 {
		return 0, &areaError{msg: "半径非法", radius: radius}
	}
	return math.Pi *  math.Pow(radius, 2), nil
	
}

func main() {
	// 自定义错误

	r, err := roundArea(2.0)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(r)

	r1, err := roundArea(-2.0)
	if err !=nil {
		fmt.Println(err)
		if err, ok := err.(*areaError);ok {
			fmt.Printf("半径是： %.2f\n", err.radius)
			fmt.Printf("msg： %s\n", err.msg)
		}
	}
	fmt.Println(r1)

}
