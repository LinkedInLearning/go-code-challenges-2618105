package bubblesort

import "fmt"

func ExampleBubbleSort() {
	s := []int{3, 2, 1, 5}
	BubbleSort(s)
	fmt.Println(s)
	// Output: [1 2 3 5]
}

func ExampleBubbleSort_second() {
	s := []int{1, 2, 3, 5}
	BubbleSort(s)
	fmt.Println(s)
	// Output: [1 2 3 5]
}
