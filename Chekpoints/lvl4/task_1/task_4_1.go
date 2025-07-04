package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

// методы реализующие интерфейс
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}
	fmt.Println(circle.Area())
	fmt.Println(rectangle.Area())
}
