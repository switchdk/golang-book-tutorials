package main

import (
	"fmt"
	"math"
)

// Shape Interface for shapes such as Circle and Rectangle
type Shape interface {
	area() float64
	perimeter() float64
}

// Circle Circle with x, y coordinates and radius r
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

// Rectangle Width and Lenght of rectangle sides, assuming symmetric rectangle
type Rectangle struct {
	l, w float64
}

func (r *Rectangle) area() float64 {
	return r.w * r.l
}

func (r *Rectangle) perimeter() float64 {
	return 2*r.l + 2*r.w
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())

	r := Rectangle{w: 2, l: 2}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
}

// 78.53981633974483
// 31.41592653589793
// 4
// 8

// Result
