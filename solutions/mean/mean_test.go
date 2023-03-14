package mean

import "fmt"

func ExampleMean() {
	input := []int{1, 2, 3, 4, 5}
	mean := Mean(input)
	fmt.Println(mean)
	// Output: 3
}

func ExampleMean_second() {
	input := []int{34, 44, 56, 78}
	mean := Mean(input)
	fmt.Println(mean)
	// Output: 53
}
