package main

import "fmt"

//append() 为切片追加元素
func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//调用append函数必须用原来的切片变量接受返回值
	//append追加元素，原来的底层数组放不下的时候，Go底层会把底层数组换一个
	//必须用变量接收append的返回值
	s1 = append(s1, "广州") 
	
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}