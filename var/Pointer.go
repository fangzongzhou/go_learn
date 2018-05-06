package main

import "fmt"

func main() {

	var a =5
	var b=5
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(*&a)
	c :=new(int)
	fmt.Println(c)
}
