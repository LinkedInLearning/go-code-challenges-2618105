package fizzbuzz

import (
	"fmt"
	"strconv"
	"strings"
)

func RenderFizzBuzzUntil(n int) string {
	var b strings.Builder
	for i := 1; i <= n; i++ {
		fmt.Fprintln(&b, translateNumberToFizzBuzz(i))
	}
	return b.String()
}

func translateNumberToFizzBuzz(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(i)
	}
}
