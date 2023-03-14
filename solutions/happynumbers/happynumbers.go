package happynumbers

func IsHappy(n int) bool {
	seen := make(map[int]bool)

	i := n
	for i != 1 && !seen[i] {
		seen[i] = true
		i = sumOfSquares(i)
	}

	return i == 1
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
