package parentheses

func isValid(s string) bool {
	opens := []rune{}

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			opens = append(opens, r)
			continue
		}

		var targetOpen rune
		switch r {
		case ')':
			targetOpen = '('
		case ']':
			targetOpen = '['
		case '}':
			targetOpen = '{'
		}

		if len(opens) == 0 || targetOpen != opens[len(opens)-1] {
			return false
		}

		opens = opens[:len(opens)-1]
	}

	return len(opens) == 0
}
