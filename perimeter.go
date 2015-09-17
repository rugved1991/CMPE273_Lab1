package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	peri() float64
}

type Rect struct {
	length, width float64
}

func (r Rect) area() float64 {
	return r.length * r.width
}

func (r Rect) peri() float64 {
	return 2 * (r.length + r.width)
}

type Circ struct {
	r float64
}

//the Square type also has an Area function and therefore, it too, implements the Shaper interface
func (c Circ) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circ) peri() float64 {
	return 2 * math.Pi * c.r
}

func main() {
	rec := Rect{length: 5, width: 3}
	fmt.Println("Details of Rectangle are: ", rec)

	var s Shape
	s = rec //equivalent to "s = Shaper(r)" since Go identifies r matches the Shaper interface
	fmt.Println("Area of Rectangle: ", s.area())
	fmt.Println()

	fmt.Println("Perimeter of Rectangle: ", s.peri())
	fmt.Println()

	cir := Circ{r: 5}
	fmt.Println("Details of Circle are: ", cir)
	s = cir //equivalent to "s = Shaper(q)
	fmt.Println("Area of Circle: ", s.area())
	fmt.Println()

	fmt.Println("Perimeter of Circle: ", s.peri())
	fmt.Println()
}