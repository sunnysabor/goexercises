package main

import "fmt"

func main() {

	//单个循环条件
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	//经典的初始/条件/后续 for 循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//不带条件的 for循环一直执行 知道break或者return
	for {
		fmt.Println("loop")
		break
	}

	//也可以使用continue跳到下一个循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
