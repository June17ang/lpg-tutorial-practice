package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLen float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{sideLen: 100.1}
	tr := triangle{
		height: 10,
		base:   2.5,
	}

	printResult("Square", sq)
	printResult("Triangle", tr)
}

func printResult(sh string, s shape) {
	fmt.Println(sh, ": ", s.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLen * sq.sideLen
}

func (tr triangle) getArea() float64 {
	return tr.base * tr.height * 0.5
}
