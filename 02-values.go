// url https://gobyexample.com/values
// Go语言包含了比如字符串、整型、浮点型、布尔型等各种不同的数据类型
// 这里给出了一些基本的使用例子
package main

import "fmt"

func main() {
	// 字符串, 可以使用 + 来连接
	fmt.Println("go" + "lang")

	// 整型以及浮点型
	fmt.Println("1+1 =", 1 + 1)
	fmt.Println("7.0/3.0 =", 7.0 / 3.0)

	// 布尔型
	fmt.Println(true&&false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
