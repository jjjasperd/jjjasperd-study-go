package main

import "fmt"

//存放元素容器
//必须制定存放的元素类型与容量（长度）
//数组【长度】是数组的【一部分】
func main() {
	var a1 [3]bool //[true false true]
	var a2 [4]bool //[true false true false]

	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//初始化方式1
	a1 = [3]bool{true, false, true}
	fmt.Println(a1)
	//初始化方式2 根据初始值自动腿短长度
	a100 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a100)

	//初始化方式3 根据索引初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	//数组遍历
	citys := [...]string{"北京", "上海", "深圳"}
	// 1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	// for range
	for i, v := range citys {
		fmt.Println(i, v)
	}

	//多维数组
	//[[1,2],[3,4],[5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)

	//多维数组遍历
	for _, v:= range a11 {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	//数组是值类型（不是引用类型 深复制）
	b1 := [3]int{1,2,3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
