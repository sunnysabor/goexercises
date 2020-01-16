package main

import "fmt"

/**
go协程在执行上来说时轻量级的线程
*/
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//假设我们有一个函数叫做 f(s) 一般这样同步 synchronously调用
	f("direct")

	//使用go f(s) 在一个Go协程中调用这个函数
	//这个新的Go协程将会并发concurrently执行这个函数
	go f("goroutin")

	//也可以为匿名函数启动一个Go协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//现在两个Go协程在独立
}
