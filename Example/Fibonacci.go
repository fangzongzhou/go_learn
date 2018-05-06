package main

import (
	"fmt"

)



func fibonacci(n int) int{
	a:=0
	b:=1
	for i:=0;i<n;i++{
		a,b=b,a+b
	}


	return a



}
func main() {
	for i:=1;i<5;i++{

		fmt.Println(fibonacci(i))
	}

}
