package longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	counts := make([]int, 26)
	maxCount, start, maxLength := 0, 0, 0

	for end := 0; end < len(s); end++ {
		counts[s[end]-'A']++
		maxCount = max(maxCount, counts[s[end]-'A'])
		if end-start+1-maxCount > k {
			counts[s[start]-'A']--
			start++
		}
		maxLength = max(maxLength, end-start+1)
	}
	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
