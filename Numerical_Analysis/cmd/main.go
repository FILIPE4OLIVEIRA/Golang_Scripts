package main

import (
	"fmt"
	"math"

	ODE "github.com/Filipe4Oliveira/numerical_analysis/internal/diferentialEquations"
	Integration "github.com/Filipe4Oliveira/numerical_analysis/internal/integration"
	Root "github.com/Filipe4Oliveira/numerical_analysis/internal/rootfinding"
)

func main() {

	// ============================================================
	// 1) ROOT FINDING (Encontrar raízes de equações f(x) = 0)
	// ============================================================

	/*
		Problema:
			f(x) = x³ - x - 2 = 0

		Solução analítica:
			Não possui forma fechada simples.

		Solução numérica esperada:
			x ≈ 1.5213797068

		Verificação (WolframAlpha):
			solve x^3 - x - 2 = 0
	*/

	f := func(x float64) float64 {
		return math.Pow(x, 3) - x - 2
	}

	// --- Método da Bisseção ---
	r1, err := Root.BisectionMethod(1, 2, 1e-6, f)
	if err != nil {
		fmt.Println("Bisection error:", err)
	} else {
		fmt.Printf("[Bisection]        root ≈ %.8f (expected ≈ 1.52138)\n", r1)
	}

	// --- Método da Secante ---
	r2, err := Root.SecantMethod(1, 2, 1e-6, f)
	if err != nil {
		fmt.Println("Secant error:", err)
	} else {
		fmt.Printf("[Secant]           root ≈ %.8f (expected ≈ 1.52138)\n", r2)
	}

	// --- Método de Newton-Raphson ---
	r3, err := Root.NewtonRaphson(1.5, 1e-6, f)
	if err != nil {
		fmt.Println("Newton-Raphson error:", err)
	} else {
		fmt.Printf("[Newton-Raphson]   root ≈ %.8f (expected ≈ 1.52138)\n\n", r3)
	}

	// ============================================================
	// 2) NUMERICAL INTEGRATION (Integração Numérica)
	// ============================================================

	/*
		Problema:
			∫sin(x)dx, x ∈ [0, π]

		Solução analítica:
			∫sin(x)dx, x ∈ [0, π] = 2

		Valor esperado:
			2.0000000000

		Verificação (WolframAlpha):
			integrate sin(x) from 0 to pi
	*/

	g := func(x float64) float64 { return math.Sin(x) }
	a, b := 0.0, math.Pi
	n := 1000

	trap, _ := Integration.TrapezoidalRule(a, b, n, g)
	fmt.Printf("[Trapezoidal]      ∫sin(x) ≈ %.10f (exact = 2)\n", trap)

	simp, _ := Integration.SimpsonComposite(a, b, n, g)
	fmt.Printf("[Simpson]          ∫sin(x) ≈ %.10f (exact = 2)\n", simp)

	mid, _ := Integration.MidpointRule(a, b, n, g)
	fmt.Printf("[Midpoint]         ∫sin(x) ≈ %.10f (exact = 2)\n\n", mid)

	/*
		Problema:
		Área entre as curvas y = sin(x) e y = cos(x) no intervalo x ∈ [0, π].

		Área entre curvas:
			Área = ∫|sin(x)-cos(x)|dx , x ∈ [0, π]

		Solução analítica:
			∫|sin(x)−cos(x)|dx, x ∈ [0, π] = 2√2 ≈ 2.8284271247

		WolframAlpha:
			integrate abs(sin(x) - cos(x)) x ∈ [0, π]

	*/

	area, _ := Integration.AreaBetweenCurves(0, math.Pi, n, math.Sin, math.Cos)
	fmt.Printf("[Area between curves] ≈ %.10f\n\n", area)

	// ============================================================
	// 3) FIRST ORDER ODE (EDO de 1ª ordem)
	// ============================================================

	/*
		Problema:
			y' = x·y
			y(0) = 1

		Solução analítica:
			y(x) = exp(x² / 2)

		Valor esperado em x = 1:
			y(1) = e^{0.5} ≈ 1.6487212707

		Verificação (WolframAlpha):
			solve y' = x y, y(0) = 1
			y(1)
	*/

	odeF := func(x, y float64) float64 {
		return x * y
	}

	_, yVec, _ := ODE.Euler(0, 1.0, 1.0, odeF)
	fmt.Printf("[Euler]            y(1) ≈ %.6f (expected ≈ 1.64872)\n",
		yVec[len(yVec)-1])

	_, yVec2, _ := ODE.RungeKutta4(0, 1.0, 1.0, odeF)
	fmt.Printf("[RK4]              y(1) ≈ %.6f (expected ≈ 1.64872)\n\n",
		yVec2[len(yVec2)-1])

	// ============================================================
	// 4) SECOND ORDER ODE (EDO de 2ª ordem)
	// ============================================================
	/*
		Problema:
			y'' + y = 0
			y(0) = 0
			y'(0) = 1

		Solução analítica:
			y(x)  = sin(x)
			y'(x) = cos(x)

		Valores esperados:
			y(π)  = 0
			y'(π) = -1

		Verificação (WolframAlpha):
			solve y'' + y = 0, y(0) = 0, y'(0) = 1
			y(pi), y'(pi)
	*/

	P := func(x float64) float64 { return 0 }
	Q := func(x float64) float64 { return 1 }
	R := func(x float64) float64 { return 0 }

	_, yVec3, dyVec3 := ODE.RungeKutta4SecondOrder(0, math.Pi, 0, 1, P, Q, R)

	fmt.Printf(
		"[RK4 2nd order]    y(π) ≈ %.6f (expected 0), y'(π) ≈ %.6f (expected -1)\n",
		yVec3[len(yVec3)-1],
		dyVec3[len(dyVec3)-1],
	)
}
