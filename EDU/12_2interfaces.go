package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}
func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}


// func main() {
// 	rectangle := Rectangle{width: 10, height: 5}
// 	circle := Circle{radius: 7}

// 	PrintShapeInfo("треугольник", rectangle)
// 	PrintShapeInfo("круг", circle)
// }

type Shaper interface {
	Area() float64
	Perimeter() float64
}

func PrintShapeInfo(title string, s Shaper) {
	fmt.Printf("Фигура %q имеет площадь %f, периметр %f\n", title, s.Area(), s.Perimeter())

}