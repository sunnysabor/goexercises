package main

import "fmt"

/**
Go协程在执行上来说是轻量级的线程
*/
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//假设有一个函数为f(s) 一般会这样同步（synchronously）调用
	f("direct")
	//使用go f(s)在一个Go协程中调用这个函数
	//这个新的Go协程将会并发（concurrently）执行这个函数
	go f("goroutine")

	//也可以为匿名函数 启动一个Go协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//现在这两个Go协程在独立的Go协程中异步（asynchronously）运行
	//程序直接运行到这一行。Scanln需要我们在程序推出前按下任意键结束
	fmt.Scanln()
	fmt.Println("done")

}
