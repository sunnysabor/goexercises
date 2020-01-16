package main

import "fmt"

//fact函数在到达 fact(0)之前一直调用自身
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

/**
支持递归 实现一个经典的阶乘示例
*/
func main() {
	fmt.Println(fact(7))
}
