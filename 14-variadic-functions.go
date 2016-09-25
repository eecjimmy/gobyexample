// url: https://gobyexample.com/variadic-functions
// 可变参数函数可以被任意个数的参数来调用
// 比如: fmt.println就是一个公共的可变参数函数

package main

import "fmt"

// 这个函数可以接受任意个数的整型参数
func sum(nums...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// 以常规的方式就可以调用到可变参数函数
	// 参数的个数完全可以自定义
	// 如果你在一个切片之中已经有了多个参数了
	// 那么只要按照func(slice...)的方式就可以调用可变参数函数了
	sum(1)
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

// Go语言中另外一个特色的功能就是闭包的功能, 下一节将来讲解这方面的内容.
