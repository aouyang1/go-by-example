package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width = 2
	return r.width * r.height
}

func (r rect) perim() int {
	r.width = 1
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	fmt.Println("rect: ", r)

	rp := &r
	rp.width = 3
	rp.height = 2
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
