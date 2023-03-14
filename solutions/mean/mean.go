package mean

func Mean(numbers []int) float64 {
	n := len(numbers)
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return float64(sum) / float64(n)
}
