package main

import "fmt"

type areaError struct {
	msg           string  // 错误的描述信息
	lenght, width float64 // 长和宽
}

func (e *areaError) Error() string {
	return e.msg
}

func (e *areaError) lenghtNegative() bool {
	return e.lenght < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(lenght, width float64) (float64, error) {
	msg := ""
	if lenght < 0 {
		msg = "长度不能小于0"
	}

	if width < 0 {
		if msg == "" {
			msg = "宽度小于零"
		} else {
			msg += ", 宽度也小于零"
		}
	}

	if msg != "" {
		return 0, &areaError{msg: msg, lenght: lenght, width: width}
	}
	return lenght * width, nil

}

func main() {
	length, width := -5.6, -7.2
	area, err := rectArea(length, width)
	if err != nil {
		fmt.Println(err)
		//if ins, ok := err.(*areaError); ok {
		//	fmt.Printf("长为： %f, 宽为： %f\n", ins.lenght, ins.width)
		//}
		if err, ok := err.(*areaError); ok {
			if err.lenghtNegative() {
				fmt.Printf("error: 长度， %.2f,小于零\n", err.lenght)
			}

			if err.widthNegative() {
				fmt.Printf("error: 宽度， %.2f, 小于零\n", err.width)
			}
		}
		return
	}
	fmt.Printf("面积为： %f\n", area)
}
