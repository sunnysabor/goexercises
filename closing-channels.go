package main

import "fmt"

//关闭一个通道意味着不能再向这个通道发送值了。这个特性可以用来给通道的接收方传达工作已经完成的消息

//使用一个jobs通道来传递main中Go协程任务执行的结束信息到一个工作Go协程中
//当我们没有多余的任务给这个工作Go协程时 我们将close这个jobs通道
func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	//这是工作Go协程  使用j,more:=<-jobs循环的从
	//jobs接收数据。在接收的这个特殊的二值形式的值中
	//如果jobs已经关闭了，并且通道中所有的值都已经接收完毕
	//那么more的值将是false.当我们完成所有的任务时，将使用这个特性通过done通道去进行通知
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	//这里使用jobs发送3个任务到工作函数中 然后关闭jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job:", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	//方法等待任务结束
	fmt.Println(<-done)
}
