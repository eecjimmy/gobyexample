// url: https://gobyexample.com/if-else
// Go语言中的if/else的分支结构是直接的
package main

import "fmt"

func main() {
	// 这是一个基础的例子
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if语句可以没有else
	if 8 % 4 == 0 {
		fmt.Println("8 is divisible of 4")
	}

	// 可以在条件之前声明变量,并且这个变量可以在所有分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// 注意: 在Go语言中, 条件无需使用(), 但是{}是必需的
	//       在Go语言中, 没有三元运算符, 因此即使是简单的条件判断, 也必需要使用完整的if结构
}
