package DiferentialEquation

func Euler_Method(xIncial, xFinal, yIncial float64,
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

func Runge_Kutta_O4(xIncial, xFinal, yIncial float64,
	MathFunction func(float64, float64) float64) (xVector, DiferentialPoint []float64, erro error) {

	var xPoint, yPoint, dydxPoint []float64
	var xNext, yNext, step float64
	var rk1, rk2, rk3, rk4 float64
	var NumberBreak int

	NumberBreak = 100000

	step = (xFinal - xIncial) / float64(NumberBreak)

	xPoint = []float64{xIncial}
	yPoint = []float64{yIncial}
	dydxPoint = []float64{0}

	for i := 0; i < NumberBreak; i++ {

		rk1 = step * MathFunction(xPoint[i], yPoint[i])
		rk2 = step * MathFunction(xPoint[i]+(step/2), yPoint[i]+rk1*(step/2))
		rk3 = step * MathFunction(xPoint[i]+(step/2), yPoint[i]+rk2*(step/2))
		rk4 = step * MathFunction(xPoint[i]+step, yPoint[i]+rk3)

		yNext = yPoint[i] + (1.0/6.0)*(rk1+2*rk2+2*rk3+rk4)
		xNext = xPoint[i] + step

		yPoint = append(yPoint, yNext)
		xPoint = append(xPoint, xNext)
		dydxPoint = append(dydxPoint, 2*xPoint[i]*yPoint[i])
	}
	return xPoint, yPoint, erro
}
