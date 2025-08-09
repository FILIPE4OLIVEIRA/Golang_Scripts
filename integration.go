package Integration

import (
	"math"
	"math/rand"
	"time"
)

func Simple_Trapezio_Method(xIncial float64, xFinal float64, NumberBreak int,
	MathFunction func(float64) float64) (Integral float64, erro error) {

	var IntegrationSum float64
	var IntegralValue float64
	var step float64

	IntegrationSum = 0.0

	step = (xFinal - xIncial) / float64(NumberBreak)

	for i := 1; i < NumberBreak; i++ {
		IntegrationSum += 2.0 * MathFunction(xIncial+float64(i)*step)
	}

	IntegralValue = (step / 2.0) * (MathFunction(xIncial) + IntegrationSum + MathFunction(xFinal))

	return IntegralValue, erro
}

func Compose_Simpson_Method(xIncial float64, xFinal float64, NumberBreak int,
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

func Average_Point_Method(xIncial float64, xFinal float64, NumberBreak int,
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

func Area_Between_2Curves(xIncial float64, xFinal float64, NumberBreak int,
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

func mean(list []float64) (average float64) {
	var sum, mean float64
	sum = 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	mean = sum / float64(len(list))
	return mean
}

func Integral_Dupla(xIncial, xFinal, yIncial, yFinal float64,
	MathFunction func(float64, float64) float64) (Integral float64, erro error) {

	var ValueIntegralAproximation []float64
	var AleatNumbers, count, MaxInt int
	var IntegrationSum, AproximationRandom, IntegralValue float64
	var xRandomPoint, yRandomPoint float64

	count = 0
	MaxInt = 1000
	AleatNumbers = 100000

	rand.Seed(time.Now().UnixNano())

	for count < MaxInt {
		IntegrationSum = 0
		for i := 0; i < AleatNumbers; i++ {
			xRandomPoint = xIncial + (xFinal-xIncial)*rand.Float64()
			yRandomPoint = yIncial + (yFinal-yIncial)*rand.Float64()
			IntegrationSum += MathFunction(xRandomPoint, yRandomPoint)
		}

		AproximationRandom = (xFinal - xIncial) * (yFinal - yIncial) * (IntegrationSum / float64(AleatNumbers))
		ValueIntegralAproximation = append(ValueIntegralAproximation, AproximationRandom)
		count = count + 1

	}
	IntegralValue = mean(ValueIntegralAproximation)
	return IntegralValue, erro
}

func Integral_Tripla(xIncial, xFinal, yIncial, yFinal, zIncial, zFinal float64,
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
	IntegralValue = mean(ValueIntegralAproximation)
	return IntegralValue, erro
}
