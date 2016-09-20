// url: https://gobyexample.com/variables
// 在Go语言中, 编译器要求必须明确指定变量类型并且变量要被使用
// 比如说会严格检查函数调用的类型是否正确
package main

import "fmt"

func main() {
	// var可以定义一个以上的变量
	var a string = "initial"
	fmt.Println(a)

	// 你可以一次性定义多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go语言可以自己推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 没有设置初始值的变量, 将会被初始化为零值. 比如int类型变量的零值是0
	var e int
	fmt.Println(e)

	// :=语法可以快速的定义和初始化一个变量, 比如说这个例子里面的相当于 var string f = "short"
	f := "short"
	fmt.Println(f)
}
