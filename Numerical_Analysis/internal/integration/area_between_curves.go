package integration

import "math"

func AreaBetweenCurves(xInicial float64, xFinal float64, NumberBreak int,
	f func(float64) float64, g func(float64) float64) (float64, error) {

	h := (xFinal - xInicial) / float64(NumberBreak)
	sum := 0.0

	for i := 0; i < NumberBreak; i++ {
		xLeft := xInicial + float64(i)*h
		xRight := xLeft + h

		diffLeft := math.Abs(f(xLeft) - g(xLeft))
		diffRight := math.Abs(f(xRight) - g(xRight))

		sum += (diffLeft + diffRight) / 2.0
	}

	area := h * sum
	return area, nil
}
