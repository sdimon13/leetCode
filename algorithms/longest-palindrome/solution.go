package longest_palindrome

func longestPalindrome(s string) int {

	myMap := make(map[rune]bool)
	max := 0
	for _, r := range s {
		if _, ok := myMap[r]; ok {
			max += 2
			delete(myMap, r)
		} else {
			myMap[r] = true
		}
	}

	if len(myMap) > 0 {
		max++
	}

	return max
}
