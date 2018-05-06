package main

import "fmt"

type Employee struct {
	id int
	name string


}

type point struct {
	x int
	y int
}

type circle struct {
	point
	radius int
}

type wheel struct {
	circle
	spokes int
}




var fzz Employee

func main() {

	var w wheel
	w.x=1
	w.y=2
	w.radius=6
	w.spokes=8
	fmt.Println(w)
	fzz.id=101
	fzz.name= "fzz"
	fzz1:=Employee{id:101,name: "fzz"}
	fmt.Println(fzz==fzz1)
	fmt.Println(fzz)
	var grr =Employee{id:2,name: "grr"}
	fmt.Println(grr)
}
