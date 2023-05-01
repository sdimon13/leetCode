package fizz_buzz

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {

		strVal := ""

		if i%3 == 0 {
			strVal += "Fizz"
		}

		if i%5 == 0 {
			strVal += "Buzz"
		}

		if strVal == "" {
			strVal += strconv.Itoa(i)
		}

		result[i-1] = strVal
	}
	return result
}

func fizzBuzzV2(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {

		strVal := ""

		if i%3 == 0 {
			strVal += "Fizz"
		}

		if i%5 == 0 {
			strVal += "Buzz"
		}

		if strVal == "" {
			strVal += strconv.Itoa(i)
		}

		result[i-1] = strVal
	}
	return result
}
