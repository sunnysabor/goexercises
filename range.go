package main

import "fmt"

func main() {

	//使用range来对slice中的元素求和
	//对于数组也可以采用这种方法
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range在数组 和slice中提供对每项的索引和值的访问
	//上面不需要索引 所以使用_来忽略它 有时候是需要这个索引的
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	//range在map中迭代键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	//range也可以只遍历map的健
	for k := range kvs {
		fmt.Println("key", k)
	}
	fmt.Println("---------")
	//range在字符串中迭代unicode码点（code point）
	//第一个返回值是字符的起始字节位置 然后第二个是字符本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
