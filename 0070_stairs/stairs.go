package stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	a := 0
	b := 1
	f := 1

	for i := 1; i < n; i++ {
		a, b = b, f
		f = a + b
	}

	return f
}
