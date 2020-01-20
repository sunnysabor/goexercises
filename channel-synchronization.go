package main

import (
	"fmt"
	"time"
)

/**
使用通道来同步Go协程间的执行状态。
这里是一个使用阻塞的接受方式来等待一个Go协程的运行结束
*/

//这是一个我们将要在Go协程中运行的函数。'done'通道将被用于通知其他Go协程这个函数已经工作完毕
func worker(done chan bool) {
	fmt.Println("working ....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	//运行一个worker Go协程，并给予用于通知的通道
	done := make(chan bool, 1)
	go worker(done)

	//程序将在接收到通道中worker 发出的通知前一直阻塞
	fmt.Println(<-done)
}
