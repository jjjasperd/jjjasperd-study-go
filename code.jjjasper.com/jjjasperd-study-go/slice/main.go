package main

import "fmt"
//切片是引用类型(值随底层数组改变)
//切片本质：框住一块数据结构一致的连续内存
//切片不能直接用于比较
func main() {
	var s1 []int //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放string类型的元素切片
	fmt.Println(s1, s2)
	
	//初始化
	s1 = []int{1,2,3}
	s2 = []string{"test1", "test2", "test3"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false
	
	//长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1),cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2),cap(s2))

	//由数组得到的切片
	a1 := [...]int{1,3,5,6,9,11,13}
	s3 := a1[0:4] //基于数组切割，左闭右开 左包含右不包含
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4] //[0:4]
	s6 := a1[3:] //[3:len(a1)]
	s7 := a1[:] //[0:len(a1)]
	fmt.Println(s5,s6,s7)
	//切片容量为底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5),cap(s5))
	// 底层数组从切片的第一个元素到最后元素的数量·
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6),cap(s6))
	//切切切切（切片再切片）
	s8 := s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8),cap(s8))

	// make函数创建切片
	ms1 := make([]int, 5, 10)
	fmt.Printf("ms1=%v len(ms1)=%d capm(ms1) = %d\n", ms1, len(ms1), cap(ms1))

	ms2 := make([]int, 0, 10)
	fmt.Printf("ms2=%v len(ms2)=%d capm(ms1) = %d\n", ms2, len(ms2), cap(ms2))
	

}