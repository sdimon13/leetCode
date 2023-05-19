package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	pCounter := make([]int, 26)
	sCounter := make([]int, 26)
	result := make([]int, 0)

	for _, ch := range p {
		pCounter[ch-'a']++
	}

	for i, ch := range s {
		sCounter[ch-'a']++
		if i >= len(p) {
			sCounter[s[i-len(p)]-'a']--
		}
		if equalCounters(pCounter, sCounter) {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

func equalCounters(counter1 []int, counter2 []int) bool {
	for i := 0; i < 26; i++ {
		if counter1[i] != counter2[i] {
			return false
		}
	}
	return true
}
