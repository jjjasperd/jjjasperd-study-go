package main

import "fmt"

func main() {
	//老生常谈基本款
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//省略初始语句
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	//省略计算
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//死循环
	// for {

	// }
	

	//for range
	s := "Hello大佬"

	for j,v := range s {
		fmt.Printf("%d %c\n", j, v)
	}

	//打印99乘法表
	for i :=1 ; i < 10; i++ {
		for j:=1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i,j,i*j)
		}
		println()
	}
}
