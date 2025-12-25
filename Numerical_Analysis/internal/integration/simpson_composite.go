package integration

func SimpsonComposite(xIncial float64, xFinal float64, NumberBreak int,
	MathFunction func(float64) float64) (Integral float64, erro error) {

	var IntegrationSum, IntegrationSum1, IntegrationSum2 float64
	var IntegralValue float64
	var step float64

	IntegrationSum1 = 0.0
	IntegrationSum2 = 0.0

	step = (xFinal - xIncial) / float64(NumberBreak)

	for i := 1; i < (NumberBreak/2)-1; i++ {
		IntegrationSum1 += 2.0 * MathFunction(xIncial+2.0*float64(i)*step)
	}

	for i := 1; i < (NumberBreak / 2); i++ {
		IntegrationSum2 += 4.0 * MathFunction(xIncial+2.0*float64(i-1)*step)
	}

	IntegrationSum = IntegrationSum1 + IntegrationSum2
	IntegralValue = (step / 3.0) * (MathFunction(xIncial) + IntegrationSum + MathFunction(xFinal))
	return IntegralValue, erro
}
