package main

import "fmt"

//当使用通道作为函数的参数时 可以指定这个通道是不是只用来发送或者接收值 这个特性提升了程序的类型安全性
//ping函数定义了一个只允许发送数据的通道。尝试使用这个通道
//来接收数据将会得到一个编译时错误
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//pong函数定义了一个只允许接收来自通道(ping)的数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println("received:", msg)
	msg = "after pong" + msg
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
