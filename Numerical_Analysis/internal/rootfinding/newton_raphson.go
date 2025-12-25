package rootfinding

import (
	"math"

	"github.com/Filipe4Oliveira/numerical_analysis/internal/suport"
)

func NewtonRaphson(xIncial float64, delta float64,
	MathFunction func(float64) float64) (ZeroPoint float64, erro error) {

	var MaxInt int
	var count int
	var valueNext float64

	MaxInt = 100
	count = 1

	if math.Abs(MathFunction(xIncial)) < delta {
		return xIncial, erro
	} else {
		for count < MaxInt && math.Abs(MathFunction(xIncial)) > delta {
			valueNext = xIncial - (MathFunction(xIncial) / suport.DyDx(xIncial, MathFunction))
			xIncial = valueNext
			count = count + 1
		}

		if count > MaxInt {
			return float64(MaxInt), erro
		} else {
			return xIncial, erro
		}
	}
}
