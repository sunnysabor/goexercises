package main

import "fmt"

//这个函数接受任意数目的int作为参数
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	//常规方式调用 传入独立的参数
	sum(1, 2)
	//如果含有一个或多个的slice 把它们作为参数使用 需要这样调用
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
