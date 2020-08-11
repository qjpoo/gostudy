package main

import (
	"fmt"
)

type Outer struct {
	a string
}

func (o *Outer) Hello(in *Inner) {
	fmt.Println("Inner:", in, "\tcall\tOuter:", o.a)
}

type Inner struct {
	a  string
	cb Cb
}

type Cb func(*Inner)

func (in *Inner) Register(cb Cb) {
	in.cb = cb
}

func (in *Inner) Say() {
	in.cb(in)
}

func main() {
	out1 := Outer{a: "out1 instance"}
	out2 := Outer{a: "out2 instance"}

	fmt.Println("out1 :", &out1)
	fmt.Println("out2 :", &out2)

	in1 := Inner{a: "int1 instance"}

	in1.Register(out1.Hello)
	in1.Say()
	in1.Register(out2.Hello)
	in1.Say()

}