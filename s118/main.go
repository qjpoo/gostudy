package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ticket = 100  // 100张票
func main() {
	// 临界资源安全问题
/*
	a := 1

	go func() {
		a = 2
		fmt.Println("goroutes中a: ", a)
	}()

	a = 3
	fmt.Println("main goroutes中a: ", a)

 */


	// 卖票
	go saleTicket("售票口1")
	go saleTicket("售票口2")
	go saleTicket("售票口3")
	go saleTicket("售票口4")

	time.Sleep(5 * time.Second)
}

func saleTicket(name string)  {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0{
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Println(name, "售出", ticket)
			ticket--
		}else {
			fmt.Println(name, "售完,没有票了...")
			break
		}
	}

}