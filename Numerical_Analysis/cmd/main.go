package main

import (
	"fmt"
	"math"

	diferentialEquation "github.com/Filipe4Oliveira/numerical_analysis/internal/ode"
)

// P(x), Q(x), R(x)
func P(x float64) float64 {
	return -2
}

func Q(x float64) float64 {
	return 2
}

func R(x float64) float64 {
	return math.Exp(2*x) * math.Sin(x)
}

func main() {

	x0 := 0.0
	x1 := 0.9

	y0 := -0.4  // y(0)
	dy0 := -0.6 // y'(0)

	x, y, dy := diferentialEquation.RungeKutta4SecondOrder(
		x0, x1,
		y0, dy0,
		P, Q, R,
	)

	fmt.Println("Runge-Kutta 4ª ordem (EDO 2ª ordem)")
	fmt.Printf("x final = %.2f\n", x[len(x)-1])
	fmt.Printf("y(x)    = %.10f\n", y[len(y)-1])
	fmt.Printf("y'(x)   = %.10f\n", dy[len(dy)-1])
}
