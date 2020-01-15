package main

import "fmt"

func main() {

	//创建空map 需要使用内建的make
	//make(map[key-type]val-type)
	m := make(map[string]int)

	//使用典型的 make[key]=val 语法来设置键值对
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	//使用name[key]来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	//当对一个map调用内建的len时 返回的时键值对数目
	fmt.Println("len", len(m))

	//内建的delete可以从一个map中移除 键值对
	delete(m, "k2")
	fmt.Println("map", m)

	//从一个map中取值时 可选的第二返回值表示这个健是否存在这个map中 可以一哦那个来消除健不存在和健有零值
	//像 0 或者 "" 而产生的歧义 这里我们不需要值 所以用_空白标示符忽略
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//可以通过这个语法在同一行声明和初始化一个新的map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)
}
