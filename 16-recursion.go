// url: https://gobyexample.com/recursion
// Go语言支持递归函数
// 下面是一个经典的阶乘的例子
package main

import "fmt"

// fact会一直调用自身直到n=0的时候才会停止
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}
func main() {
	fmt.Println(fact(7))
}
