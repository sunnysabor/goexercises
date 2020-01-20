package main

import "fmt"

//默认情况下 通道是没有缓冲的 意味着只有对应的接收（<- chan） 通道准备好接收时，才允许进行发送
//（chane <-）.可缓冲通道允许在没有对应接收方的情况下，缓存限定数量的值
func main() {

	//这里我们创建了一个字符串通道 最多允许缓存2个值
	messages := make(chan string, 2)
	//由于通道是缓冲的 可以将这些值发送到通道中
	//而不需要相应的并发接收
	messages <- "buffered"
	messages <- "channel"

	//接收这两个值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
