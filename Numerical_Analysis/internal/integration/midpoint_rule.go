package integration

func MidpointRule(xInicial float64, xFinal float64, NumberBreak int,
	MathFunction func(float64) float64) (float64, error) {

	h := (xFinal - xInicial) / float64(NumberBreak)
	sum := 0.0

	for i := 0; i < NumberBreak; i++ {
		mid := xInicial + (float64(i)+0.5)*h
		sum += MathFunction(mid)
	}

	integral := h * sum
	return integral, nil
}
