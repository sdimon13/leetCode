package merge_strings_alternately

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		res += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	if i < len(word1) {
		res += word1[i:]
	}

	if j < len(word2) {
		res += word2[j:]
	}

	return res
}
