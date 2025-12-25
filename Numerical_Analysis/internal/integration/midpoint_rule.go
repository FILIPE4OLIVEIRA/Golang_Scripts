package integration

func MidpointRule(xIncial float64, xFinal float64, NumberBreak int,
	MathFunction func(float64) float64) (Integral float64, erro error) {

	var IntegrationSum float64
	var IntegralValue float64
	var step float64

	IntegrationSum = 0.0

	step = (xFinal - xIncial) / float64(NumberBreak)

	for i := 1; i < (NumberBreak / 2); i++ {
		IntegrationSum += 2.0 * MathFunction(xIncial+float64(i+1)*step)
	}

	IntegralValue = (2.0 * step) * IntegrationSum

	return IntegralValue, erro
}
