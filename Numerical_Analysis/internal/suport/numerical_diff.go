package suport

func DyDx(xIncial float64, MathFunction func(float64) float64) (tetha float64) {
	var step float64 = 0.001
	tetha = (MathFunction(xIncial+step) - MathFunction(xIncial)) / (step)
	return tetha
}
