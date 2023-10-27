package main

import "fmt"

// 2.2 常量的声明
// const PI = 3.14

// func main() {
// 	fmt.Println(PI)
// }

//2.4 常量类型 - 类型常量
// const A int = 1 // 使用已定义的类型 int

// func main() {
// 	fmt.Println(A)
// }

//2.4 常量类型 - 无类型常量
// const A = 1 // 没有已知的类型，而是使用类型推断

// func main() {
// 	fmt.Println(A)
// }

// 2.5 常量的不可更改且只读
// func main() {
// 	const A = 1
// 	A = 2 // VS code 提示报错
// 	fmt.Println(A)
// }

// 2.6 多常量声明
// 多常量声明
const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
