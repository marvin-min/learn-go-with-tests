package struct_method_interface

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}

func (triangle Triangle) Area() float64 {
	return (triangle.Base * triangle.Height) / 2
}
