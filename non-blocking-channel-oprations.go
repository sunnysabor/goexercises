package main

import "fmt"

/**
常规通过通道发送和接收数据是阻塞的
我们可以使用带default的子句的select来实现非阻塞的发送接收 甚至是非阻塞的多路select
*/
func main() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	default:
		fmt.Println("no message received")
	}

	//一个非阻塞发送的实现方式和上面一样
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//可以在default前使用多个case子句来实现一个多路的非阻塞的选择器。
	//试图在messages和signals上同时使用非阻塞的接收操作
	select {
	case msg := <-messages:
		fmt.Println("received:", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
