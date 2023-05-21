package backspace_string_compare

func backspaceCompare(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1

	for i >= 0 || j >= 0 {
		i = getNextIndex(s, i)
		j = getNextIndex(t, j)

		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}

		if (i >= 0) != (j >= 0) {
			return false
		}

		i--
		j--
	}
	return true
}

func getNextIndex(str string, index int) int {
	skip := 0
	i := index
	for i >= 0 {
		if str[i] == '#' {
			skip++
			i--
		} else if skip > 0 {
			skip--
			i--
		} else {
			break
		}
	}
	return i
}
