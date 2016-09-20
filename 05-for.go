// url: https://gobyexample.com/for
// for是Go语言中仅有的循环结构, 下面给出了三种基本的循环类型
package main

import "fmt"

func main() {
	// 最常用的类型, 使用单条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化/条件/后续操作结构
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 没有条件语句的for将会一直循环直到break或者在函数中return
	for {
		fmt.Println("loop")
		break
	}

	// 之后我们将会在range声明/通道或者其它的数据结构中看到for的其它一些用法
}
