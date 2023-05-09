package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[rune]rune)
	tToS := make(map[rune]bool)

	for index, sRune := range s {
		tRune := rune(t[index])

		if mappedTRune, exists := sToT[sRune]; exists {
			if mappedTRune != tRune {
				return false
			}
		} else {
			if tToS[tRune] {
				return false
			}
			sToT[sRune] = tRune
			tToS[tRune] = true
		}
	}
	return true
}

func isIsomorphicV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := [256]int{}
	tToS := [256]int{}

	for i := 0; i < len(s); i++ {
		sChar, tChar := s[i], t[i]

		if sToT[sChar] != tToS[tChar] {
			return false
		}

		sToT[sChar] = i + 1
		tToS[tChar] = i + 1
	}

	return true
}
