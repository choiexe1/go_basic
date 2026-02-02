package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (c *Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t *Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func PrintInfo(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Printf("%f\n", s.Area())
	fmt.Printf("%f\n", s.Perimeter())
}
