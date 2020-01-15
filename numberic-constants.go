package main

import (
	"fmt"
	"math/cmplx"
)

const(
	A bool =true
	B int=1
	C string ="hello world"
	D float32=1.22222
	Big =1<<100
	Small=Big>>99

)

func needInt(x int)  {
	var a int =1
	fmt.Println(a)
}

func main() {
	 e :=cmplx.Sqrt(-5+12i)
	fmt.Println(e)

}
