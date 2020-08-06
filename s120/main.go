package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 20  // 100张票
var mutex sync.Mutex  // 申请一个锁的对象
var wg sync.WaitGroup
func main() {
	// mutx 互斥锁
	// 加锁之后,一定要解锁, 建议用defer mutex.unclock()

	// 卖票
	wg.Add(4)
	go saleTicket("售票口1")
	go saleTicket("售票口2")
	go saleTicket("售票口3")
	go saleTicket("售票口4")
	wg.Wait()
	//time.Sleep(5 * time.Second)
	fmt.Println("main end ...")
}

func saleTicket(name string)  {
	rand.Seed(time.Now().UnixNano())
	for {
		// 31-41之间的代码,只能有一个goroutine来执行,相当于没有并发了,相当于单线程了, g1上了锁,别的就只能等他解锁了,才能运行
		mutex.Lock()
		if ticket > 0{
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Println(name, "售出", ticket)
			ticket--
		}else {
			mutex.Unlock()
			fmt.Println(name, "售完,没有票了...")
			break
		}
		mutex.Unlock()
	}
	//wg.Done()
	defer wg.Done()

}