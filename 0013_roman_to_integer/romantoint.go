package romantoint

var twos = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

var ones = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	num := 0

	for i := 0; i < len(s); i++ {
		c := rune(s[i])

		if i < len(s)-1 {
			sub := s[i : i+2]
			if inc, ok := twos[sub]; ok {
				num += inc
				i++
				continue
			}
		}

		inc, _ := ones[c]
        num += inc
	}

	return num
}
