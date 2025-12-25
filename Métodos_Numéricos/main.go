package main

import (
	"NumericalMethods/Integration"
	screen "fmt"
	"math"
)

func MathEquation(x, y, z float64) (point float64) {
	result := math.Sin(x * y * z)
	return result
}

func main() {
	var x0, x1, y0, y1, z0, z1 float64

	x0 = 0.1
	x1 = 0.9
	y0 = 0.2
	y1 = 0.5
	z0 = 0.4
	z1 = 0.7

	Integral, erro := Integration.Integral_Tripla(x0, x1, y0, y1, z0, z1, MathEquation)

	if erro != nil {
		screen.Printf("Integral_Dupla \n")
		screen.Printf("Could not Itegrate in This Interview! \n")
	} else {
		screen.Printf("Integral_Dupla \n")
		screen.Printf("The Integral is: %.15v \n", Integral)
	}

}
