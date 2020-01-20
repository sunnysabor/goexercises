package main

import (
	"fmt"
	"time"
)

/**
超时对于一个链接外部资源，或者其他一些需要花费执行时间的操作程序而言是很重要的
得益于通道和select,在Go中实现超时操作是简洁优雅的
*/
func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	//这里使用select实现一个超时操作
	//"res:=<-c1"等待结果,<-Time.After 等待超时
	//时间1秒后发送的值。由于select默认处理第一个已准备好的接收操作
	//如果这个操作超过了允许的1秒 将会执行超时case
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	//如果我允许一个长一点的超时时间3秒，将会成功从c2 接收到值并且打印结果
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
