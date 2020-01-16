package main

import "fmt"

//通过两个函数 zeroval和zeroptr来比较指针和值类型的不同 zeroval有一个int型参数所以使用值传递
//zeroval将从调用它的那个函数中得到一个ival形参的拷贝
func zeroval(ival int) {
	ival = 2
}

//zeroptr有和上面不同的*int参数 意味着它用了一个int指针。
//函数体内的*iptr接着 ——解引用——这个指针 从它的内存地址得到这个地址对应的当前值。对一个解引用的指针赋值将会改变这个指针引用的真实地址的值
func zeroptr(iptr *int) {
	*iptr = 2
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	//通过&i语法来取得i的内存地址 即指向i的指针
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	//指针也是可以被打印的
	fmt.Println("pointer:", &i)
}
