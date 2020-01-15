package main

import (
	"fmt"
	"time"
)

func main() {

	//基本姿势
	i := 2
	fmt.Println("write ", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	}
	//同一个case中 可以使用逗号来分隔多个表达式 下面使用了可选的default分支
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("workday")
	}

	//不带表达式的switch是实现if/else逻辑的另一种
	//case也可以不使用常量
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its after noon")
	}

	//类型开关 比较类型而非值。可以用来发现一个接口值的类型 下面 t在每个分支中会有相应的类型
	var whatAmI = func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			bytes, err := fmt.Printf("Don't know type %T\n", t)
			fmt.Println(bytes, err)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
