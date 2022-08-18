package main

import (
	"fmt"
	"github.com/vz250049/fluffy-eureka/internal"
	"github.com/vz250049/fluffy-eureka/pkg/circle"
	"github.com/vz250049/fluffy-eureka/pkg/rectangle"
)

func main() {
	var s internal.Shape = circle.Circle{Radius: 5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	s = rectangle.Rectangle{Length: 4.0, Width: 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())
}
