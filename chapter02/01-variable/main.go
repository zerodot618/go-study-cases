package main

import "fmt"

// 1.1.4 变量声明并初始化
// func main() {
// 	var student1 string = "Haha" // type 类型是字符串
// 	var student2 = "Dudu"        // type 使用类型推断
// 	x := 2                       // type 也是使用类型推断

// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)
// }

// 1.1.5 变量声明未初始化
// func main() {
// 	var a string
// 	var b int
// 	var c bool

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// 1.1.6 变量先声明后赋值
// func main() {
// 	var student1 string
// 	student1 = "Haha"
// 	fmt.Println(student1)
// }

// 1.7 var 和 := 的区别 - var的示例
// var a int     // 可以函数外声明
// var b int = 2 // 可以函数外声明并赋值
// var c = 3     // 可以函数外声明并类型推断赋值

// func main() {
// 	a = 1 // 可在函数内赋值
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// 1.7 var 和 := 的区别 - var的示例
// a := 1 // 这行编辑器会提示错误，所以 := 是不能在函数外使用的

// func main() {
//   fmt.Println(a)
// }

// 1.2.1 多变量声明 - 同一行中声明多个变量
// func main() {
// 	var a, b, c, d int = 1, 3, 5, 7

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// 1.2.1 多变量声明 - 同一行中声明不同类型的变量
// func main() {
// 	var a, b = 6, "Hello"
// 	c, d := 7, "World!"

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// 1.2.2 代码块多变量声明
func main() {
	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
