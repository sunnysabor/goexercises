package main

import "fmt"

//这个intSeq 函数返回另一个在 intSeq函数体内定义的匿名函数。这个返回的函数使用闭包的方式 _隐藏_变量 i
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	//调用intSeq函数 将返回值（一个函数）赋给nextInt。这个函数的值包含了自己的值i，这样在每次调用nextInt时都会更新i的值
	nextInt := intSeq()

	//通过多次调用nextInt来看看闭包的效果
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//为了确认这个状态对于特定的函数是唯一的 重新创建并测试一下
	newInts := intSeq()
	fmt.Println(newInts())
}
