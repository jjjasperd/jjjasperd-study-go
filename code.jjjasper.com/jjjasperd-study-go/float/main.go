package main

import (
	"fmt"
	"math")


func main() {
	 
	fmt.Println(math.MaxFloat32) //float32 最大值 
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go小数都是float64
	f2 := float32(1.2345)
	fmt.Printf("%T\n", f2) //默认go小数都是float64
	// float32 不能直接复制给 float64
	
}