package main

import (
	"fmt"
	"tutorial2/internal"
	"tutorial2/pkg/circle"
	"tutorial2/pkg/rectangle"
)

func main() {
	var s internal.Shape = circle.Circle{Radius: 5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	s = rectangle.Rectangle{Length: 4.0, Width: 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())
}
