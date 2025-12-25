package suport

func Mean(list []float64) float64 {
	var sum float64
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum / float64(len(list))
}
