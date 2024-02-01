package sqrt

func mySqrt(x int) int {
	sqr := func(x int) int {
		return x * x
	}

	lo := 0
	hi := x
	mid := (lo + hi) / 2

	for lo <= hi {
		switch {
		case sqr(mid) == x:
			return mid

		case sqr(mid) > x:
			hi = mid
			mid = (lo + hi) / 2

		case sqr(mid) < x && sqr(mid+1) > x:
			return mid

		case sqr(mid) < x && sqr(mid+1) < x:
			lo = mid
			mid = (lo + hi) / 2

		default:
			return mid + 1
		}
	}

	return -1
}
