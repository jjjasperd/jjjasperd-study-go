package main

import (
	"fmt";
	"strings"
)


func main() {
	s := "Gogogo"
	fmt.Println(s)

	s2 := `
		我要
			涨
				工资
	`
	fmt.Println(s2)

	s3 := `sdasd:{a\s\d}`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))
	
	// 字符串拼接
	name := "hsj"
	world :="msn"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf("%s%s\n", name, world)
	fmt.Println(ss1)

	//分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	//判断包含
	fmt.Println(strings.Contains(ss, "asdasd"))
	fmt.Println(strings.Contains(ss, "msn"))

	//前缀、后缀
	fmt.Println(strings.HasPrefix(ss, "hsj"))
	fmt.Println(strings.HasSuffix(ss, "hsj"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	//拼接
	fmt.Println(strings.Join(ret,"+"))
}