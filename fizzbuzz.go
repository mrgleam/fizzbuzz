package fizzbuzz

import "strconv"

func fizzbuzz(number int) string {
	var result string
	if number%3 == 0 {
		result = "fizz"
	}
	if number%5 == 0 {
		result = result + "buzz"
	}
	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}
