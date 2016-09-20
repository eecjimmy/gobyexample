// url: https://gobyexample.com/slices
// Slices是Go语言中的主要数据类型
// 提供了一个比数组更为强大的接口序列
package main

import "fmt"

func main() {
	// 原文: Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// 使用内置的make方法可以创建一个非零长度的空切片
	// 这里我们生成了一个长度为3的字符串类型的切片(初始为零值)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 可以像数组一样进行赋值和取值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len返回切片的长度
	fmt.Println("len:", len(s))

	// 除了那些内置操作之外, 切片另外还提供了几个比数组更丰富多彩的操作
	// 其中一个就是append, 这个内置方法将会增加一个或者多个值到切片中, 并返回这个切片
	// 注意: 必须要使用一个切片来接收append返回的新的切片
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 切片也可以被复制
	// 这里我们创建了一个和s相同长度的切片c, 并且把s复制给c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片还提供了一个slice[low:high]的语法来进行"切"操作
	// 比如下面的操作, 将获取s[2], s[3], s[4]这三个元素
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个操作将从开始"切"到s[5](不包括s[5])
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个操作将从s[2]"切"到结束(包括s[2])
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们同样可以在单行内为定义和初始化一个切片变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 可以组成多维的切片数据
	// 不像数组, 切片的内部切片的长度是可以改变的
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j;
		}
	}
	fmt.Println("2d: ", twoD)

	// 原文: Note that while slices are different types than arrays, they are rendered similarly by fmt.Println.
	// 查看Go团队写的的关于设计和实现切片的一篇不错的博文: http://blog.golang.org/2011/01/go-slices-usage-and-internals.html
	// 现在我们已经介绍了数组和切片了, 下面将会介绍map
}
