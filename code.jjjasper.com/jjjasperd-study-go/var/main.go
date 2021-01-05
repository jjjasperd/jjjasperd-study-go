package main

import "fmt"

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "搞钱"
	age = 29
	isOk = true
	//Go 语言非全局变量声明必须使用，不然编译不过
	var s1 string = "Old wang"
	//类型推导
	var s2 = "20"
	//简短变量声明,只能在函数中使用
	s3 := "Short"
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
