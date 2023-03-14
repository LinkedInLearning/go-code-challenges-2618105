package prime

import "fmt"

func ExampleFindPrimeNumbersInInterval() {
	primes := FindPrimeNumbersInInterval(1, 10)
	fmt.Println(primes)
	// Output: [2 3 5 7]
}
