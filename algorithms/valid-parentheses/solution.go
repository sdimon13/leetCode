package valid_parentheses

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []rune

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
