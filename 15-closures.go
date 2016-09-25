// url: https://gobyexample.com/closures
// Go语言支持匿名函数
// 匿名函数在你想要在行内定义一个无名函数的时候, 非常有用.

package main

import "fmt"

// 我们在intSeq的函数体中定义了一个函数
// 然后intSeq将会返回这个函数
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// 我们使用nextInt来接收intSeq的返回值(一个函数)
	nextInt := intSeq()

	// 当我们调用nextInt的时候, 将会每次捕获到i的值
	// 通过多调用几次, 我们可以看到效果
	fmt.Println(nextInt()) // print 1
	fmt.Println(nextInt()) // print 2
	fmt.Println(nextInt()) // print 3

	// 我们来创建一个新的intSeq来验证他们的i是不会冲突的
	newInts := intSeq()
	fmt.Println(newInts()) // print 1
}
// 最后来看下函数的最后一个特性 -- 递归.