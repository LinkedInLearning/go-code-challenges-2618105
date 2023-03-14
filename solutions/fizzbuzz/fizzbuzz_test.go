package fizzbuzz

import "fmt"

func ExampleRenderFizzBuzzUntil() {
	output := RenderFizzBuzzUntil(15)
	fmt.Println(output)
	// Output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}