package main

import "fmt"

func main() {
	//基本姿势
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 si odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//在条件语句之前可以有一个声明语句 这里声明的变量可以用在所有的条件分支中
	a:=8
	if num := a; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
