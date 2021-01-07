package main

import "fmt"

func main() {
	//bool in go cannot convert to other type.
	b1 := true
	var b2 bool
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value: %v \n", b2, b2)
}