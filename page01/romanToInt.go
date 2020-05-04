package page01

func RomanToInt(s string) int {
	dict := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
		'#': 0,
	}

	sum, length := 0, len(s)
	arr := []rune(s)
	for i, c := range arr {
		if i < length-1 {
			switch {
			case c == 'I' && arr[i+1] == 'V':
				sum = sum + 3
				arr[i+1] = '#'
			case c == 'I' && arr[i+1] == 'X':
				sum = sum + 8
				arr[i+1] = '#'
			case c == 'X' && arr[i+1] == 'L':
				sum = sum + 30
				arr[i+1] = '#'
			case c == 'X' && arr[i+1] == 'C':
				sum = sum + 80
				arr[i+1] = '#'
			case c == 'C' && arr[i+1] == 'D':
				sum = sum + 300
				arr[i+1] = '#'
			case c == 'C' && arr[i+1] == 'M':
				sum = sum + 800
				arr[i+1] = '#'
			}
		}
		sum = sum + dict[c]
	}

	return sum
}
