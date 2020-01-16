package main

import "fmt"

//(int,int)在这个函数中标志这这个函数返回2个int
func vals() (int, int) {
	return 3, 7
}

func main() {

	//可以通过-多赋值-操作 来使用这两个不同的返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//如果仅仅需要返回值的一部分的话 可以使用空白标示符 _
	_, c := vals()
	fmt.Println(c)
}
