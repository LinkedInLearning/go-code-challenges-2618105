package bubblesort

func BubbleSort(input []int) {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
}
