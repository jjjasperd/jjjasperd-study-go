package main

import "fmt"

func main() {
	s := "Hello小强"
	fmt.Printf("%s\n", s)
	
	//字符串修改 (和java一样 =immutable)
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换成一个rune切片
	s3[0] = '红' //char int32 单引号
	fmt.Println(string(s3))// 把 rune 切片强制转换成字符串
	c := byte('H') //byte(uint8)
	fmt.Printf("%T\n", c)

	//类型转换
	n := 10
	var f float64
	f = float64(n)
	fmt.Println(f)
}