// url: https://gobyexample.com/functions
// 函数是Go语言的核心.
// 下面将会通过一些不同的例子来学习Go语言中的函数
package main

import "fmt"

// 这是一个函数
// 接受两个int类型参数
// 返回int类型的和
// Go语言中需要显示返回
// 比如说他不能自动返回函数的最后一条表达式
func plus(a int, b int) int {
	return a + b
}

// 当你有多个连续类型的参数时
// 除了最后一个类型不可以省略, 前面的相同类型的参数是可以省略的
func plusPlus(a, b, c int) int {
	return a + b + c;
}
func main() {
	// 你可以使用name(args)的方式调用你想调用的函数
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}

// 在Go语言中函数还有其他几个功能.
// 其中一个就是多个返回值, 将在下一节中讲解
