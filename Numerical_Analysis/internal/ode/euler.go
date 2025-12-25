package diferentialEquation

func Euler(xIncial, xFinal, yIncial float64,
	MathFunction func(float64, float64) float64) (xVector, DiferentialPoint []float64, erro error) {

	var NumberBreak int
	var xPoint, yPoint []float64
	var xNext, yNext, step float64

	NumberBreak = 100000

	step = (xFinal - xIncial) / float64(NumberBreak)

	xPoint = []float64{xIncial}
	yPoint = []float64{yIncial}

	for i := 0; i < NumberBreak; i++ {

		yNext = yPoint[i] + step*MathFunction(xPoint[i], yPoint[i])
		xNext = xPoint[i] + step

		yPoint = append(yPoint, yNext)
		xPoint = append(xPoint, xNext)
	}
	return xPoint, yPoint, erro
}
