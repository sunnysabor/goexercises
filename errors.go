package main

import (
	"errors"
	"fmt"
)

/**
符合go语言习惯的做法是 使用一个独立 明确的返回值来传递错误信息。go语言的处理方式能清楚知道哪个函数返回了错误
并能像调用那些没有出错的函数一样调用
*/

//按照惯例 错误通常是最后一个返回值并且是error类型 一个内建的接口
func f1(arg int) (int, error) {
	if arg == 42 {
		//errors.New构造一个使用给定的错误信息的基本error值
		return -1, errors.New("can not work with 42")
	}
	//返回错误值为 nil代表没有错误
	return arg + 3, nil
}

//通过实现Error方法来自定义error类型是可以的
//使用自定义错误类型来表示上面例子中的参数错误
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//这个例子中 使用&argError语法来建立一个新的结构体
		//提供了arg和prob这两个字段的值
		return -1, &argError{arg, "can not work with it"}
	}
	return arg + 3, nil
}

func main() {

	//下面的两个循环测试了返回错误的函数 注意在if行内的错误检查代码 在Go中是一个普遍的用法
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}

	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("F2 FAILED:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	//如果想再程序中使用一个自定义错误类型中的数据 需要通过类型断言来得到这个错误类型的实例
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
