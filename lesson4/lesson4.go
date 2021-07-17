package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	height float64
	width  float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
func (s Circle) String() string {
	return fmt.Sprintf("Cirlce: radius %v", s.radius)
}
func (s Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %v and width %v", s.height, s.width)
}

func (a Rectangle) Area() float64 {
	return a.height * a.width
}
func (a Circle) Area() float64 {
	return math.Pi * math.Pow(a.radius, 2)
}
func (a Rectangle) Perimeter() float64 {
	return 2 * (a.height + a.width)
}
func (a Circle) Perimeter() float64 {
	return 2 * math.Pi * a.radius
}
func main() {

	c := Circle{radius: 8}
	r := Rectangle{
		height: 9,
		width:  3,
	}
	DescribeShape(c)
	DescribeShape(r)
}
