package main

import "fmt"

func main() {
	var student1 string = "Haha" // type 类型是字符串
	var student2 = "Dudu"        // type 使用类型推断
	x := 2                       // type 也是使用类型推断

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
