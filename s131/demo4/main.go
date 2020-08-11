package main

import "fmt"

type A struct {
	Name string
}

func (a A) Foo(x, y int)  {
	fmt.Println("x: ", x, "y: ", y, a.Name)

}
func main() {
	b := new(A)
	b.Name = "chiling"
	foob := b.Foo
	foob(1,2)
	b.Foo(11,22)


}
