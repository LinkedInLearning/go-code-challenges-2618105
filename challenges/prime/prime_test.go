package prime

import "fmt"

func ExampleFindPrimes() {
	primes := FindPrimes(1, 10)
	fmt.Println(primes)
	// Output: [2 3 5 7]
}
