package main

import "fmt"

//这里person结构体包含了name和age两个字段
type person struct {
	name string
	age  int
}

func main() {
	//创建新的结构体元素
	fmt.Println(person{"Bob", 20})
	//初始化一个结构体元素时指定字段名字
	fmt.Println(person{name: "Alice", age: 30})

	//省略的字段将被初始化成为零值
	fmt.Println(person{name: "Fred"})

	//& 前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 40})

	//使用.来访问结构体字段
	s := person{name: "Sean", age: 50}
	fmt.Println(s.age)

	//结构体是可变的
	s.age = 51
	fmt.Println(s.age)
}
