// url: https://gobyexample.com/switch
// 多个分支可以使用switch进行判断
package main

import (
	"fmt"
	"time"
)

func main() {
	// 下面是一个最简单的switch结构
	i := 2
	fmt.Print("write ", i, " as ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 多个条件在同一个case中, 可以使用,隔开
	// 这个例子中的我们使用了可选的default情况
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// 没有表达式的switch其实是另外一种if/else的表达方式
	// 这里展示了case表达式中如何使用不确定值
	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}
