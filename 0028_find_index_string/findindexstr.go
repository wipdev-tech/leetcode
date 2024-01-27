package findindexstring

func strStr(haystack string, needle string) int {
	n := len(needle)
	for i := range haystack {
		if i+n > len(haystack) {
			break
		}

		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
