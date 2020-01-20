package main

import "fmt"

/**
通道是链接多个Go协程的管道。可以从一个Go协程将值发送到通道 然后在别的Go协程中接收
*/
func main() {

	//使用make(chan  val-type)创建一个新的通道
	//通道类型就是需要传递值的类型
	messages := make(chan string)
	//使用channel <- 语法 _发送(send)_一个新的值到通道中。这里我们在一个新的Go协程中发送"ping"到上面创建的
	//messages通道中
	go func() { messages <- "ping" }()
	go func() { messages <- "test " }()
	//使用<-channel语法从通道中 _接收（receives）_一个值。这里将接收我们在上面发送的ping消息并打印出来
	msg := <-messages
	fmt.Println(msg)
}
