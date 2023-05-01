package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	var counts [26]int

	if len(ransomNote) > len(magazine) {
		return false
	}

	for _, l := range magazine {
		counts[l-'a']++
	}

	for _, l := range ransomNote {
		if counts[l-'a'] == 0 {
			return false
		} else {
			counts[l-'a']--
		}
	}
	return true
}
