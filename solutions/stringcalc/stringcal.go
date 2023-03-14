package stringcalc

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	if numbers == "" {
		return 0
	} else {
		split := strings.Split(numbers, ",")
		sum := 0
		for _, number := range split {
			i, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			sum += i
		}
		return sum
	}
}
