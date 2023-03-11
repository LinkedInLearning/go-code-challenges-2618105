package prime

func FindPrimes(min, max int) []int {
	var primes []int

	for i := min; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
