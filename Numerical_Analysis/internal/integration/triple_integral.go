package integration

import (
	"math/rand"
	"time"

	"github.com/Filipe4Oliveira/numerical_analysis/internal/suport"
)

func TripleIntegral(xIncial, xFinal, yIncial, yFinal, zIncial, zFinal float64,
	MathFunction func(float64, float64, float64) float64) (Integral float64, erro error) {

	var ValueIntegralAproximation []float64
	var AleatNumbers, count, MaxInt int
	var IntegrationSum, AproximationRandom, IntegralValue float64
	var xRandomPoint, yRandomPoint, zRandomPoint float64

	count = 0
	MaxInt = 1000
	AleatNumbers = 100000

	rand.Seed(time.Now().UnixNano())

	for count < MaxInt {
		IntegrationSum = 0
		for i := 0; i < AleatNumbers; i++ {
			xRandomPoint = xIncial + (xFinal-xIncial)*rand.Float64()
			yRandomPoint = yIncial + (yFinal-yIncial)*rand.Float64()
			zRandomPoint = zIncial + (zFinal-zIncial)*rand.Float64()
			IntegrationSum += MathFunction(xRandomPoint, yRandomPoint, zRandomPoint)
		}

		AproximationRandom = (xFinal - xIncial) * (yFinal - yIncial) * (zFinal - zIncial) * (IntegrationSum / float64(AleatNumbers))
		ValueIntegralAproximation = append(ValueIntegralAproximation, AproximationRandom)
		count = count + 1

	}
	IntegralValue = suport.Mean(ValueIntegralAproximation)
	return IntegralValue, erro
}
