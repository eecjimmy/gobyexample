// url: https://gobyexample.com/maps
// maps是Go语言中内置的关联数组(有时候在其他语言中也叫做散列或者字典)
package main

import "fmt"

func main() {
	// 可以使用内置方法make创建一个空的map:
	// make(map[key-type][val-type])
	m := make(map[string]int)

	// 可以使用经典的name[key]的语法来设置一对map键值对
	m["k1"] = 7
	m["k2"] = 13

	// 在使用一个类似Println的方法打印map时
	// 将会显示map的完整键值对
	fmt.Println(m)

	// 可以使用name[key]的方式来获取一个键的值
	v1 := m["k1"]
	fmt.Println(v1)

	//当在map上使用len方法时, 将会返回map的键值对的个数
	fmt.Println("len:", len(m))

	// 内置方法delete可以从map中移除一个键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 第二个返回值用于判断map中是否存在这个key
	// 这通常用来消除零值和键值不存在所产生的歧义
	// 在这个例子中, 我们没有用到value
	// 因此用空白标识符_来忽略
	_, prs := m["k5"]
	fmt.Println(prs)

	// 也可以用以下语法在单行里面对map进行定义和赋值
	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n)

	// 注意: 当使用fmt.Println进行打印的时候
	// 将会按照[k:v k:v]的形式输出
}
