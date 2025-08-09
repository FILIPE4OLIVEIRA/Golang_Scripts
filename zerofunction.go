package ZeroFunction

import (
	"fmt"
	"math"
)

func Bisection_Method(xIncial float64, xFinal float64, delta float64,
	MathFunction func(float64) float64) (ZeroPoint float64, erro error) {

	var count int
	var average float64
	var valueInicial float64
	var valueFinal float64
	var valueAverage float64

	count = 1
	valueInicial = MathFunction(xIncial)
	valueFinal = MathFunction(xFinal)

	if valueInicial*valueFinal < 0 {
		for math.Abs(xFinal-xIncial) > delta {
			average = (xFinal + xIncial) / 2.0
			valueAverage = MathFunction(average)
			if math.Abs(xIncial-xFinal) < delta || math.Abs(valueAverage) < delta {
				break
			} else {
				if valueInicial*valueAverage > 0 {
					xIncial = average
				} else {
					xFinal = average
				}
			}
			count = count + 1
		}
		return average, erro
	} else {
		erro = fmt.Errorf("%d", 1)
		return average, erro
	}
}

func dydx(xIncial float64, MathFunction func(float64) float64) (tetha float64) {
	var step float64 = 0.001
	tetha = (MathFunction(xIncial+step) - MathFunction(xIncial)) / (step)
	return tetha
}

func Newton_Raphson(xIncial float64, delta float64,
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
			valueNext = xIncial - (MathFunction(xIncial) / dydx(xIncial, MathFunction))
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

func Secante_Method(xIncial float64, xFinal float64, delta float64,
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
