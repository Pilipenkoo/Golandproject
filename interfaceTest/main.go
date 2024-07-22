package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	Height float64
	width  float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) area() float64 {
	return r.Height * r.width
}

func (c Circle) area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func main() {
	rect := Rectangle{Height: 10, width: 20}
	fmt.Println(rect.area())
	circle := Circle{Radius: 5}
	fmt.Println(circle.area())
}
