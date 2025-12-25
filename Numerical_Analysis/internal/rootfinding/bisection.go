package rootfinding

import (
	"fmt"
	"math"
)

func BisectionMethod(xIncial float64, xFinal float64, delta float64,
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
