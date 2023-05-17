package climbing_stairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2

	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return b
}
