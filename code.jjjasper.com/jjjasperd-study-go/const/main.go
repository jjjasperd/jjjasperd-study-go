package main

import "fmt"

const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时， 如果某一行声明后没有复制，默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota const出现，从0起，每增加一行未声明变量就加一
const (
	iota1 = iota //0
	iota2 //1
	iota3 //2
)

const (
	b1 = iota //0
	b2 = iota //1
	_  = iota//2 隐式 占位
	b3 = iota//3
)

const (
	c1 = iota //0
	c2 = 100 //100 被覆盖 
	c3 = iota//2 每新增一行加一
	c4 = iota//3
)

const (
	d1, d2 = iota +1, iota + 2 //d1:1 d2:2
	d3, d4 = iota +1, iota + 2 // d3: 2, d4:3
)

//定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("iota1:", iota1)
	fmt.Println("iota2:", iota2)
	fmt.Println("iota3:", iota3)

	
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
}