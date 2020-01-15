package main

import "fmt"

/**
类型推导
 */
func main() {
	v :=42.123
	s :="ss"
	var a int
	b :=a
	fmt.Printf("v is of type %T\n",v)
	fmt.Printf("v is of type %T\n",s)
	fmt.Printf("v is of type %T\n",b)

}
