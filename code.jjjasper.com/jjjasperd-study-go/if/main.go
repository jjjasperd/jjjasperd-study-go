package main

import "fmt"

func main() {
	age := 10
	if age > 18 {
		fmt.Println("是时候赚钱养家了")
	} else {
		fmt.Println("放下手机，滚去写作业")
	}
	//if特殊写法
	//作用域
	//a 只在if判断中生效
	if a:=19; a>18 {
		fmt.Println("是时候赚钱养家了")
	} else {
		fmt.Println("放下手机，滚去写作业")
	}
}