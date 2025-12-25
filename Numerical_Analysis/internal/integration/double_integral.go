package integration

import (
	"math/rand"
	"time"

	"github.com/Filipe4Oliveira/numerical_analysis/internal/suport"
)

func DoubleIntegral(xIncial, xFinal, yIncial, yFinal float64,
	MathFunction func(float64, float64) float64) (Integral float64, erro error) {

	var ValueIntegralAproximation []float64
	var AleatNumbers, count, MaxInt int
	var IntegrationSum, AproximationRandom, IntegralValue float64
	var xRandomPoint, yRandomPoint float64

	count = 0
	MaxInt = 1000
	AleatNumbers = 100000

	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for count < MaxInt {
		IntegrationSum = 0
		for i := 0; i < AleatNumbers; i++ {
			xRandomPoint = xIncial + (xFinal-xIncial)*localRand.Float64()
			yRandomPoint = yIncial + (yFinal-yIncial)*localRand.Float64()
			IntegrationSum += MathFunction(xRandomPoint, yRandomPoint)
		}

		AproximationRandom = (xFinal - xIncial) * (yFinal - yIncial) * (IntegrationSum / float64(AleatNumbers))
		ValueIntegralAproximation = append(ValueIntegralAproximation, AproximationRandom)
		count = count + 1

	}
	IntegralValue = suport.Mean(ValueIntegralAproximation)
	return IntegralValue, erro
}
