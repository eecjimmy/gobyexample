// url: https://gobyexample.com/range
// range可以在各种数据结构中进行遍历
// 下面让我们看看在我们之前学习过的数据结构中使用range
package main

import (
	"fmt"
)

func main() {
	// 这里我们使用range对切片中的数字进行相加
	// 数组操作也是这样
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range可以在数组和切片中同时为每个条目提供键和值
	// 上面的例子我们不需要索引, 所以用空白符_来忽略
	// 可是有时候我们真正需要的是键
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 用range在map上遍历时可以返回键值对
	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 当在字符串上使用range遍历的时候, 将会返回unicode的代码点(code points)
	// 第一个值将返回字符(rune)的索引
	// 第二个值将返回字符(rune)本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
