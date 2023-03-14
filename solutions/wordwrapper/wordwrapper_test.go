package wordwrapper

import "fmt"

func ExampleWrap() {
	fmt.Println(Wrap("This is a very long sentence that will be wrapped", 10))
	// Output:
	// This is a
	// very long
	// sentence
	// that will
	// be wrapped
}
