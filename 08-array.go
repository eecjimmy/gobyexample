// url: https://gobyexample.com/arrays
// 在Go语言中, 数组是一个指定元素长度的有序队列
package main

import "fmt"

func main() {
	// 这里创建了一个长度为5的int类型的数组
	// 元素的类型以及长度都是数组类型的一部分
	// 默认的数组为置为零值, 如int就是0
	var a[5]int
	fmt.Println("emp:", a)

	// 可以使用array[index]=value的方式给数组元素赋值
	// 也可以使用array[index]的方式获取元素的值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 内置函数len返回数组的长度
	fmt.Println("len:", len(a))

	// 使用下面的语法可以定义并且初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 数组类型是一维的
	// 但是可以组成多维的数据结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 注意:
	//       当使用fmt.Println打印数组时, 将会打印出[v1,v2,v3...]这样的形式
	// 在Go语言中, slices通常比array更具有代表性, 下面就会讲到slices
}
