package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器表示在未来某一时刻的独立事件。告诉定时器需要等待的时间
	//然后它将提供一个用于通知的通道 这里定时器将等待2秒
	time1 := time.NewTimer(time.Second * 2)
	<-time1.C
	fmt.Println("Timer 1 expired")

	//如果仅仅需要单纯的等待，需要使用time.Sleep
	//定时器使用原因之一就是可以在定时器失效之前，取消这个定时器 如下
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}
