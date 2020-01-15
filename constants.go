package main

import (
	"fmt"
	"math"
)

/**
常量
**/
const Pi = 3.14

const s = "constant"

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	const n = 500
	//当上下文需要时 比如变量赋值或者函数调用
	//一个数可以被给定一个类型 例如这里的函数 需要一个 float64 的参数
	fmt.Println(math.Sin(n))
}
