package rootfinding

import (
	"fmt"
	"math"
)

func SecantMethod(xIncial float64, xFinal float64, delta float64,
	MathFunction func(float64) float64) (ZeroPoint float64, erro error) {

	var count int
	var MaxInt int
	var valueNext float64

	MaxInt = 100
	count = 1

	if MathFunction(xIncial)*MathFunction(xFinal) < 0 {

		valueNext = (xIncial + xFinal) / 2.0

		for count < MaxInt && math.Abs(MathFunction(valueNext)) > delta {
			valueNext = xIncial - ((MathFunction(xIncial) * (xIncial - xFinal)) / (MathFunction(xIncial) - MathFunction(xFinal)))
			xIncial = xFinal
			xFinal = valueNext
			count = count + 1
		}
		if count > MaxInt {
			return float64(MaxInt), erro
		} else {
			return valueNext, erro
		}
	} else {
		erro = fmt.Errorf("%d", 1)
		return valueNext, erro
	}
}
