package main

import "fmt"

func main() {
	switch n := 2; n {
	//可以多条件
	case 1, 3:
		fmt.Println("11111")
	case 2:
		fmt.Println("22222")
	default:
		fmt.Println("ddddd")
	}

}
