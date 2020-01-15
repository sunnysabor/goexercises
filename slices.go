package main

import "fmt"

func main() {

	//与数组不同 slice的类型仅由它所包含的元素决定（不需要元素的个数）
	//要创建一个长度非零的空slice 需要使用内建的方法'make' 这里我们创建了一个长度为3的string类型 slice(初始化为零值)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	//可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	//除了基本操作 slice具备丰富的。其中一个是内建的append 返回一个包含了一个或者多个新值的slice
	//由于append可能返回新的slice 需要接受其返回值
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	//slice也可以被copy 这里我们创建一个空的和s有相同长度的slice c 并且将s复制给c
	c := make([]string, len(s))
	fmt.Println("copy before", c)
	copy(c, s)
	fmt.Println("copy:", c)

	//slice支持通过slice[low:high] 语法进行切片操作 这里得到一个包含元素 2 3 4 的slice
	l := s[2:5]
	fmt.Println("sl1:", l)

	//这个slice从0切片 到5 不包含
	l = s[:5]
	fmt.Println("s12:", l)

	//可以在一行代码中声明并初始化一个slice变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//slice可以组成多位数据结构。内部的slice长度可以不一致 这和多维数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {

			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
