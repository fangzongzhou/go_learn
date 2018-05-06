package main

import "fmt"

func gcd(int1 int,int2 int) int  {
	for int2!= 0 {
		int1,int2=int2,int1%int2
	}

	return int1


}
func main() {

	fmt.Println(gcd(20,12))
}
