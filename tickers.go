package main

import "time"

/**
定时器是想在未来某一时刻执行一次使用的
打点器则是当想要在固定的时间间隔重复执行准备的
它将定时的执行 直到我们将它停止
*/
func main() {

	//打点器 和定时器的机制有点相似  一个通道用来发送数据
	//我们在这个通道使用内置的range来迭代值
	//每隔500ms发送一次的值
	ticker := time.NewTimer(time.Microsecond * 5000)

}
