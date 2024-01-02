package commonprefix

func longestCommonPrefix(strs []string) string {
	prefix := ""

	maxLen := len(strs[0])
	for _, s := range strs {
		if len(s) < maxLen {
			maxLen = len(s)
		}
	}

	for i := 0; i < maxLen; i++ {
		char := strs[0][i]

		for _, s := range strs {
			if s[i] != char {
				return prefix
			}
		}

		prefix += string(char)
	}

	return prefix
}
