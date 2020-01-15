package main

import "fmt"

func main() {

	//创建一个数组a 来存放刚好5个 int
	//元素的类型和长度都是数组类型的一部分 数组默认是零值的
	//对于int 数组来说 也就是 0
	var a [5]    int
	fmt.Println("emp:", a)

	//可以用 array[index]=value 语法来设置数组指定位置的值
	//或者用array[index]得到值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get", a[4])
	//使用内置函数len 返回数组长度
	fmt.Println("len:", len(a))

	//使用这个语法一行内声明并初始化一个数组
	b := [5]    int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//数组类型是一维的 但是可以组合 构造多位的数据结构
	var twoD [2][3] int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d: ",twoD)

}
