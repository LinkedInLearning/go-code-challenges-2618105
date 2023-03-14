package happynumbers

import "fmt"

func ExampleIsHappy() {
	result := IsHappy(19)
	fmt.Println(result)
	// Output: true
}

func ExampleIsHappy_second() {
	result := IsHappy(4)
	fmt.Println(result)
	// Output: false
}
