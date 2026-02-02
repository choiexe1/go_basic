package main

import (
	shape "go_basic/cmd/09_interfaces"
)

func main() {
	c := &shape.Circle{Radius: 5}
	r := &shape.Rectangle{Width: 3, Height: 4}
	t := &shape.Triangle{A: 3, B: 4, C: 5}

	shape.PrintInfo(c)
	shape.PrintInfo(r)
	shape.PrintInfo(t)
}
