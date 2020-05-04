package page01

func Reverse(x int) int {
	var a = x
	if x < 0 {
		a = (-1) * x
	}

	m, sum := 0, 0
	for {
		a, m = a/10, a%10
		sum = sum*10 + m

		if sum > 0x7fffffff {
			return 0
		}
		if a == 0 {
			break
		}
	}
	if x > 0 {
		return sum
	}
	return (-1) * sum
}
