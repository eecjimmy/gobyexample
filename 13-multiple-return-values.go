// url: https://gobyexample.com/multiple-return-values
// Go语言内置支持多个返回值
// 这个特性在Go语言中经常被使用
// 比如一个函数一次性返回结果和错误
package main

import "fmt"

// 这里函数签名中的(int, int)表示这个函数将会返回两个int类型
func vals() (int, int) {
	return 3, 7
}

func main() {
	// 这里我们通过调用多个返回值的函数来返回两个不同的值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 如果你仅仅想返回值中的某一个部分
	// 可以使用空白符[_]来忽略其他的值
	_, c := vals()
	fmt.Println(c)
}

// 接受数量可变的参数是Go语言中另外一个不错的功能
// 下一节将看下这方面的内容
