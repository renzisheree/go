package main

import "fmt"

type square struct{
	sideLength float64
}
type triangle struct{
	height float64
	base float64
}
type shape interface {
	getArea() float64
}

func main() {
	sq := square{sideLength: 10}
	t := triangle{height: 10 , base: 15}
	sq.getArea()
	t.getArea()
	printArea(sq)
	printArea(t)
}
func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
func (t triangle) getArea() float64 {
	return 0.5*t.base*t.height
}
func  printArea(s shape) {
	fmt.Println(s.getArea())
}