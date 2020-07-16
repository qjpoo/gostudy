package main

import "fmt"

func	f1()(int, float64, string) {
	return 10, 10.0, "hello"
}

/*
func f2(age int) int {
	if age >=0 {
		return age   // 这里也会报错, 因为if不确定会被执行,如果没有被执行的话,函数就不会有返回值了. 所以在else加一个return
	}else {
		return 0
    }
}
 */

func f3(i int) {
	for j:=1; j< i;j++ {
		if j == 5 {
			return
		}else{
			fmt.Println(j)
		}
	}

}
func main() {
	/*
	注意点:
	1. 一个函数定义了返回值, 必须使用return语句将结果返回给调用处, return后的数据必须和函数定义的个数,类型,顺序一致
	2. 结束函数的执行,那怕后面还有代码,都不会执行
	3. 返回多个值的时候,可以用_来舍弃
	4. 一个函数没有定义返回值,也是可以使用return来结束函数的
	*/
	i, f, s := f1() //如果不想要的话,可以使用空白标识符 _ 来丢弃
	fmt.Println(i, f, s)
	f3(8)



}
