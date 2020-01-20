package main

import (
	"fmt"
	"time"
)

func main() {

	//我们的例子中 将从两个通道中选择
	c1 := make(chan string)
	c2 := make(chan string)
	//各个通道将在若干时间后接收一个值 这个用来模拟例如并行的Go协程中阻塞的RPC操作
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "two"
	}()

	//使用select关键字来同时等待这两个值 并打印各自接收到的值
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)

		}
	}
}
