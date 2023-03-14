package csvparser

import "fmt"

func ExampleParse() {
	turingAwardWinners := Parse("turing.csv")
	fmt.Println(turingAwardWinners[1968])
	// Output:
	// Richard Hamming
}

func ExampleParse_multiple() {
	turingAwardWinners := Parse("turing.csv")
	fmt.Println(turingAwardWinners[2020])
	// Output:
	// Alfred Aho,Jeffrey Ullman
}
