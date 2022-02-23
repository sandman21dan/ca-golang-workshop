package roman

var charMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func ToDecimal(roman string) int {
	// Convert the roman string to a slice of integers
	n := []int{}
	for _, c := range roman {
		// Use string() to convert from rune to string
		n = append(n, charMap[string(c)])
	}

	// start acumulator at zero
	r := 0

	for i, v := range n {
		// if the next value is greater than the current value,
		// it means we need to substract the current value
		if i+1 < len(n) && v < n[i+1] {
			r -= v
		} else {
			r += v
		}
	}

	return r
}
