package decode_string

type stackItem struct {
	str   []rune
	count int
}

func decodeString(s string) string {
	repeatCount := 0
	currentSubstring := []rune{}
	var stack []stackItem

	for _, ch := range s {
		if '0' <= ch && ch <= '9' {
			repeatCount = repeatCount*10 + int(ch-'0')
		} else if ch == '[' {
			stack = append(stack, stackItem{
				str:   currentSubstring,
				count: repeatCount,
			})
			currentSubstring = []rune{}
			repeatCount = 0
		} else if ch == ']' {
			lastStackItem := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for i := 0; i < lastStackItem.count; i++ {
				lastStackItem.str = append(lastStackItem.str, currentSubstring...)
			}
			currentSubstring = lastStackItem.str
		} else {
			currentSubstring = append(currentSubstring, ch)
		}
	}

	return string(currentSubstring)
}
