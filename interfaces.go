package main

import (
	"fmt"
	"math"
)

/**
接口时命名了的方法签名的集合
*/

//几何体基本接口
type geometry interface {
	area() float64
	perim() float64
}

//类型rects 和circle实现这个接口
type rects struct {
	width, height float64
}
type circle struct {
	radius float64
}

//要在Go实现一个接口 就需要实现接口中的所有方法
func (r rects) area() float64 {
	return r.width * r.height
}
func (r rects) perim() float64 {
	return 2*r.width + 2*r.height
}

//circle实现
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//如果一个变量具有接口类型 可以用调用指定接口中的方法
//下面是一个通用的measure函数 利用它来在任何的geometry中工作
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rects{width: 3.12, height: 4.12}
	c := circle{radius: 5.12}

	//结构体circle和rect都实现了geometry接口
	//可以使用它们的实力作为measuer的参数
	measure(r)
	fmt.Println(c)
	measure(c)
}
