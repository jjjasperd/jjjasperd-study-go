package main

import "fmt"

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) //10进制转换2进制
	fmt.Printf("%o\n", i1) //10进制转换8进制 (文件权限)
	fmt.Printf("%x\n", i1) //10进制转换16进制 (内存地址)

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	//查看变量类型
	fmt.Printf("%T\n", i3)

	i4 := int8(9) //明确指定int8类型， 否则默认int类型
	fmt.Printf("%T\n", i4)
}