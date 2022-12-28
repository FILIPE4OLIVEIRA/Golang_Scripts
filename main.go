package main

import (
	screen "fmt"
	"math"
	"metodosnumericos/Numerical"
)

func MathEquation(x float64) (point float64) {
	yx := math.Exp(-3*x) * math.Sin(4*x)
	return yx
}

func main() {
	var x0, x1, delta float64

	x0 = -0.5
	x1 = 0.5
	delta = 0.00001

	ZeroPoint1, erro1 := Numerical.Bisection_Method(x0, x1, delta, MathEquation)
	ZeroPoint2, erro2 := Numerical.Newton_Raphson(x0, delta, MathEquation)

	if erro1 != nil {
		screen.Printf("Bisection_Method \n")
		screen.Printf("There is no root in this interview! \n")
	} else {
		screen.Printf("Bisection_Method \n")
		screen.Printf("The root is: %.15v \n", ZeroPoint1)
	}

	if erro2 != nil {
		screen.Printf("Newton_Raphson \n")
		screen.Printf("There is no root in this interview! \n")
	} else {
		screen.Printf("Newton_Raphson \n")
		screen.Printf("The root is: %.15v \n", ZeroPoint2)
	}

}
