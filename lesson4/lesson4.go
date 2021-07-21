package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	height float64
	width  float64
}

func DescribeShape(s Shape) {
	var (
		v float64
		e error
	)
	v, e = s.Area()
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(s)
		fmt.Printf("Area: %.2f\n", v)
		v, _ = s.Perimeter()
		fmt.Printf("Perimeter: %.2f\n", v)
	}
}
func (c Circle) String() string {
	return fmt.Sprintf("Cirlce: radius %v", c.radius)
}
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %v and width %v", r.height, r.width)
}

func (r Rectangle) Area() (float64, error) {
	if (r.height <= 0) || (r.width <= 0) {
		return 0, errors.New("cant calculate the values for rectangle")
	}
	return r.height * r.width, nil
}
func (c Circle) Area() (float64, error) {
	if c.radius <= 0 {
		return 0, errors.New("cant calculate the values for circle")
	}
	return math.Pi * math.Pow(c.radius, 2), nil
}
func (r Rectangle) Perimeter() (float64, error) {
	if (r.height <= 0) || (r.width <= 0) {
		return 0, errors.New("cant calculate the values for rectangle")
	}
	return 2 * (r.height + r.width), nil
}
func (c Circle) Perimeter() (float64, error) {
	if c.radius <= 0 {
		return 0, errors.New("cant calculate the values  for circle")
	}
	return 2 * math.Pi * c.radius, nil
}
func main() {

	c := Circle{radius: 10}
	r := Rectangle{
		height: 9,
		width:  3,
	}
	DescribeShape(c)
	DescribeShape(r)
}
