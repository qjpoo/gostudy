package main

import "fmt"

type animal interface {
	speak()
}

type cat struct {
	name string
}

type dog struct {
	name string
}

func (c cat) speak()  {
	fmt.Println(c.name, "cat ...")
}

/*
func (c *cat) speak()  {
	fmt.Println(" cat ...")
}
*/
func (d dog) speak()  {
	fmt.Println("dog ...")
}


func main() {
	// interface

	var x animal
	cat1 := cat{name: "小猫"}
	x = cat1  // 当接受者如果是地址的话,这里会不行
	x.speak()

	cat2 := &cat{name: "小猫咪"}
	x = cat2
	x.speak() // (*cat2).speak()
	(*cat2).speak()




}
