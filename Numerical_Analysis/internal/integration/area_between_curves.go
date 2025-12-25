package integration

import (
	"math"
)

func AreaBetweenCurves(xIncial float64, xFinal float64, NumberBreak int,
	MathFunction1 func(float64) float64, MathFunction2 func(float64) float64) (Integral float64, erro error) {

	var IntegrationSum1, IntegrationSum2 float64
	var ValueIntegralArea1, ValueIntegralArea2 float64
	var ValueIntegralArea float64
	var step float64

	IntegrationSum1 = 0.0
	IntegrationSum2 = 0.0
	ValueIntegralArea = 0.0

	step = (xFinal - xIncial) / float64(NumberBreak)

	for i := 1; i < NumberBreak; i++ {
		IntegrationSum1 += 2.0 * MathFunction1(xIncial+float64(i)*step)
		IntegrationSum2 += 2.0 * MathFunction2(xIncial+float64(i)*step)
	}

	ValueIntegralArea1 = (step / 2.0) * (MathFunction1(xIncial) + IntegrationSum1 + MathFunction1(xFinal))
	ValueIntegralArea2 = (step / 2.0) * (MathFunction2(xIncial) + IntegrationSum2 + MathFunction2(xFinal))
	ValueIntegralArea = math.Abs(ValueIntegralArea1 - ValueIntegralArea2)

	return ValueIntegralArea, erro
}
