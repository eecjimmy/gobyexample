// url: https://gobyexample.com/constants
// Go语言支持定义[字符/字符串/布尔值/数字]作为常量
package main

import (
	"fmt"
	"math"
)

// const 定义一个常量值
const s string = "constant"

func main() {
	fmt.Println(s)

	// 通过var定义变量的地方就可以使用const来定义常量
	const n = 500000000

	// 常量表达式可以进行任意精读的运算
	const d = 3e20 / n
	fmt.Println(d)

	// 原文: A numeric constant has no type until it’s given one, such as by an explicit cast.
	fmt.Println(int64(d))

	// 数字也可以根据上下文自动转换类型, 比如说math.Sin方法期望的参数是float64的类型
	fmt.Println(math.Sin(n));

}
