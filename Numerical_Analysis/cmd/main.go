package main

import (
	"fmt"
	"math"

	Integration "github.com/Filipe4Oliveira/numerical_analysis/internal/integration"
	Root "github.com/Filipe4Oliveira/numerical_analysis/internal/rootfinding"
)

func main() {
	// --- Root finding examples ---
	f := func(x float64) float64 { return math.Pow(x, 3) - x - 2 } // root ~ 1.521

	r1, err := Root.BisectionMethod(1, 2, 1e-6, f)
	if err != nil {
		fmt.Println("Bisection error:", err)
	} else {
		fmt.Printf("Bisection root: %.8f\n", r1)
	}

	r2, err := Root.SecantMethod(1, 2, 1e-6, f)
	if err != nil {
		fmt.Println("Secant error:", err)
	} else {
		fmt.Printf("Secant root: %.8f\n", r2)
	}

	r3, err := Root.NewtonRaphson(1.5, 1e-6, f)
	if err != nil {
		fmt.Println("NewtonRaphson error:", err)
	} else {
		fmt.Printf("Newton-Raphson root: %.8f\n", r3)
	}

	// --- Integration examples ---
	g := func(x float64) float64 { return math.Sin(x) }
	a, b := 0.0, math.Pi

	trap, _ := Integration.TrapezoidalRule(a, b, 1000, g)
	fmt.Printf("Trapezoidal (sin) ≈ %.10f (exact = 2)\n", trap)

	simp, _ := Integration.SimpsonComposite(a, b, 1000, g)
	fmt.Printf("Simpson composite (sin) ≈ %.10f\n", simp)

	mid, _ := Integration.MidpointRule(a, b, 1000, g)
	fmt.Printf("Midpoint rule (sin) ≈ %.10f\n", mid)

	// Area between curves: sin(x) and cos(x)
	area, _ := Integration.AreaBetweenCurves(0, math.Pi, 1000, math.Sin, math.Cos)
	fmt.Printf("Area between sin and cos on [0,pi] ≈ %.10f\n", area)

	// Note: DoubleIntegral/TripleIntegral use Monte Carlo with large internal defaults and may be slow.

	// --- Differential equations ---
	// Example ODE: y' = x * y, y(0)=1 -> solution y = exp(x^2/2)
	odeF := func(x, y float64) float64 { return x * y }
	xVec, yVec, _ := ODE.Euler(0, 1.0, 1.0, odeF)
	fmt.Printf("Euler last point: x=%.6f y=%.6f\n", xVec[len(xVec)-1], yVec[len(yVec)-1])

	xVec2, yVec2, _ := ODE.RungeKutta4(0, 1.0, 1.0, odeF)
	fmt.Printf("RungeKutta4 last point: x=%.6f y=%.6f\n", xVec2[len(xVec2)-1], yVec2[len(yVec2)-1])

	// Second order example: y'' + y = 0  (simple harmonic oscillator)
	P := func(x float64) float64 { return 0 }
	Q := func(x float64) float64 { return 1 }
	R := func(x float64) float64 { return 0 }
	xVec3, yVec3, dyVec3 := ODE.RungeKutta4SecondOrder(0, math.Pi, 0, 1, P, Q, R)
	fmt.Printf("RK4 2nd order last point: x=%.6f y=%.6f dy=%.6f (expected y(pi)≈0)\n", xVec3[len(xVec3)-1], yVec3[len(yVec3)-1], dyVec3[len(dyVec3)-1])
}
