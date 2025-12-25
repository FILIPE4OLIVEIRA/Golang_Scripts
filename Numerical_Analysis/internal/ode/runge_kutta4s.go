package diferentialEquation

func RungeKutta4SecondOrder(xInicial, xFinal float64, y0, dy0 float64,
	P, Q, R func(float64) float64) (xVec, yVec, dyVec []float64) {

	const N = 50000

	h := (xFinal - xInicial) / float64(N)

	xVec = []float64{xInicial}
	yVec = []float64{y0}
	dyVec = []float64{dy0}

	for i := 0; i < N; i++ {

		x := xVec[i]
		y := yVec[i]
		dy := dyVec[i]

		// k1
		k11 := h * dy
		k12 := h * (R(x) - P(x)*dy - Q(x)*y)

		// k2
		k21 := h * (dy + 0.5*k12)
		k22 := h * (R(x+h/2) - P(x)*(dy+0.5*k12) - Q(x)*(y+0.5*k11))

		// k3
		k31 := h * (dy + 0.5*k22)
		k32 := h * (R(x+h/2) - P(x)*(dy+0.5*k22) - Q(x)*(y+0.5*k21))

		// k4
		k41 := h * (dy + k32)
		k42 := h * (R(x+h) - P(x)*(dy+k32) - Q(x)*(y+k31))

		yNext := y + (1.0/6.0)*(k11+2*k21+2*k31+k41)
		dyNext := dy + (1.0/6.0)*(k12+2*k22+2*k32+k42)
		xNext := x + h

		xVec = append(xVec, xNext)
		yVec = append(yVec, yNext)
		dyVec = append(dyVec, dyNext)
	}
	return
}
